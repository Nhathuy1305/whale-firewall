package whale_firewall
<<<<<<< Updated upstream
=======

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
>>>>>>> Stashed changes
