package install

import (
	"os"
	"path"

	"github.com/DylanGuedes/laser/files"
	"github.com/DylanGuedes/laser/state"
)

func Handle(laser *state.L) *state.L {
	createScopeFolder(laser)
	createCfgFile(laser)
	createSyncFile(laser)
	return laser
}

func createScopeFolder(laser *state.L) {
	if err := os.Mkdir(laser.Scope, os.ModePerm); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
}

func createCfgFile(laser *state.L) {
	if err := os.Mkdir(path.Join(laser.Scope, files.DefautSyncedFolder), os.ModePerm); err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}
}

func createSyncFile(laser *state.L) {
	laser.StoreState()
}
