package whale_firewall

import (
	_ "embed"
	"go.uber.org/zap"
	"sync"
)

const (
	dbCommands = `
PRAGMA foreign_keys = true;
PRAGMA busy_timeout = 1000;
PRAGMA journal_mode = WAL;
`
	dummyID      = "dummy_id"
	dummyName    = "dummy_name"
	enabledLabel = "whalefirewall.enabled"
	rulesLabel   = "whalefirewall.rules"
)

//go:embed database/schema.sql
var dbSchema string

// TODO
type RuleManager struct {
	wg       sync.WaitGroup
	stopping chan struct{}
	done     chan struct{}

	logger *zap.Logger

	newDockerClient   dockerClientCreator
	newFirewallClient firewallClientCreator

	containerTracker *container
}

type dockerClientCreator func() (dockerClient, error)

type firewallClientCreator func() (firewallClient, error)
