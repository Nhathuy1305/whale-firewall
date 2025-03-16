package whale_firewall

import (
	_ "embed"
	"github.com/docker/docker/api/types"
	"go.uber.org/zap"
	"sync"
	"whalefirewall/m/container"
	"whalefirewall/m/database"
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

type RuleManager struct {
	wg       sync.WaitGroup
	stopping chan struct{}
	done     chan struct{}

	logger *zap.Logger

	newDockerClient   dockerClientCreator
	newFirewallClient firewallClientCreator

	containerTracker *container.Tracker

	createCh chan containerDetails
	deleteCh chan string

	db        database.DB
	dockerCli dockerClient
}

type dockerClientCreator func() (dockerClient, error)

type firewallClientCreator func() (firewallClient, error)

type containerDetails struct {
	container types.ContainerJSON
	isNew     bool
}
