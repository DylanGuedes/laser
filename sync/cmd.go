package sync

import (
	kingpin "github.com/alecthomas/kingpin/v2"
)

func RegisterCmd(app *kingpin.Application) *kingpin.CmdClause {
	return app.Command("sync", `Sync current state based on current configuration files.`)
}
