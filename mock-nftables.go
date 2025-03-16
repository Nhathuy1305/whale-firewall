package whale_firewall

import (
	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"github.com/mitchellh/copystructure"
	"go.uber.org/zap"
	"sync"
)

const anonSetName = "__set%d"

var setAllocNum = 1

type firewallClient interface {
	AddTable(t *nftables.Table) *nftables.Table

	AddChain(c *nftables.Chain) *nftables.Chain
	DelChain(c *nftables.Chain)
	ListChainOfTableFamily(family nftables.TableFamily) ([]*nftables.Chain, error)

	AddSet(s *nftables.Set, vals []nftables.SetElement) error
	DelSet(s *nftables.Set)
	SetAddElements(s *nftables.Set, vals []nftables.SetElement) error
	SetDeleteElements(s *nftables.Set, vals []nftables.SetElement) error

	AddRule(r *nftables.Rule) *nftables.Rule
	DelRule(r *nftables.Rule) error
	InsertRule(r *nftables.Rule) *nftables.Rule
	GetRules(t *nftables.Table, c *nftables.Chain) ([]*nftables.Rule, error)

	Flush() error
}

type mockFirewall struct {
	logger *zap.SugaredLogger

	changed bool

	tables map[string]*table
	chains map[string]chain

	unsetLookupExprs []*expr.Lookup

	flushErr error

	bf baseFirewallReaderWriter
}

type table struct {
	Sets setMap

	newAnonSets map[string]bool
}

type setMap map[string][]nftables.SetElement

type chain struct {
	Chain *nftables.Chain
	Rules []*nftables.Rule
}

type baseFirewallReaderWriter interface {
	readBaseFirewall(f func(base *mockFirewall))
	writeBaseFirewall(f func(base *mockFirewall))
}

type mockFirewallCreatorI interface {
	newMockFirewall() *mockFirewall
	baseFirewallReaderWriter
}

func newMockFirewallCreator(logger *zap.Logger) mockFirewallCreatorI {
	m := &mockFirewallCreator{
		baseFirewall: &mockFirewall{
			tables: make(map[string]*table),
			chains: make(map[string]chain),
		},
		logger: logger,
	}

	return m
}

type mockFirewallCreator struct {
	baseFirewall *mockFirewall
	mtx          sync.RWMutex
	logger       *zap.Logger
}

func (m *mockFirewallCreator) newMockFirewall() *mockFirewall {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	newFirewall := &mockFirewall{
		logger: m.logger.Sugar(),
		tables: clone(m.baseFirewall.tables),
		chains: clone(m.baseFirewall.chains),
		bf:     m,
	}
	initTables(newFirewall)

	return newFirewall
}

func initTables(m *mockFirewall) {
	for _, t := range m.tables {
		t.newAnonSets = make(map[string]bool)
	}
}

func (m *mockFirewallCreator) readBaseFirewall(f func(base *mockFirewall)) {
	m.mtx.RLock()
	f(m.baseFirewall)
	m.mtx.RUnlock()
}

func (m *mockFirewallCreator) writeBaseFirewall(f func(base *mockFirewall)) {
	m.mtx.Lock()
	f(m.baseFirewall)
	m.mtx.Unlock()
}

func (m *mockFirewall) AddTable(t *nftables.Table) *nftables.Table {
	m.changed = true

	if _, ok := m.tables[t.Name]; !ok {
		m.tables[t.Name] = &table{
			Sets:        make(setMap),
			newAnonSets: make(map[string]bool),
		}
	}

	return t
}

func (m *mockFirewall) AddChain(c *nftables.Chain) *nftables.Chain {
	m.changed = true

	if _, ok := m.chains[c.Name]; !ok {
		m.chains[c.Name] = chain{
			Chain: c,
		}
	}

	return c
}

func clone[T any](t T) T {
	//nolint: forcetypeassert
	return copystructure.Must(copystructure.Copy(t)).(T)
}
