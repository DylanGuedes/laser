package main

import (
	"os"

	"github.com/DylanGuedes/laser/files"
	"github.com/DylanGuedes/laser/install"
	"github.com/DylanGuedes/laser/push"
	"github.com/DylanGuedes/laser/state"
	"github.com/DylanGuedes/laser/sync"
	kingpin "github.com/alecthomas/kingpin/v2"
	"golang.org/x/exp/slog"
)

var (
	app        = kingpin.New("laser", "A fancy configuration manager.").Version("0.1.0")
	pushCmd    = push.RegisterCmd(app)
	installCmd = install.RegisterCmd(app)
	syncCmd    = sync.RegisterCmd(app)
)

func main() {
	slog.New(slog.NewTextHandler(os.Stdout))
	scope := app.Flag("scope", "Scope to execute Laser.").Default(files.DefaultScope).String()
	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	enforceInstall(cmd != installCmd.FullCommand(), *scope)

	laser := state.LoadState(*scope)
	switch cmd {
	case installCmd.FullCommand():
		laser = install.Handle(laser)
	case pushCmd.FullCommand():
		laser = pushCmd.Handle(laser)
	case syncCmd.FullCommand():
		laser = sync.Handle(laser)
	}
}

func enforceInstall(enforce bool, scope string) {
	if enforce {
		if !files.Available(scope) {
			slog.Error("Laser installation not found. Please, first install Laser through './laser install' command.")
			os.Exit(1)
		}
	}
}
