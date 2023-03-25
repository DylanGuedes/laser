package install

import (
	kingpin "github.com/alecthomas/kingpin/v2"
)

func RegisterCmd(app *kingpin.Application) *kingpin.CmdClause {
	installCmd := app.Command("install", `Install Laser on the current repository.
The 'install' command will generate default Laser files on the current repository.

Usage:
	laser install

This will create the following files:
	.laser/config.yaml
	.laser/data.yaml`)

	return installCmd
}
