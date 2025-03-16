package whale_firewall

import (
	"github.com/google/nftables"
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
