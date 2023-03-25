package sync

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/DylanGuedes/laser/files"
	"github.com/DylanGuedes/laser/state"
	"golang.org/x/exp/slog"
)

// if cmd = sync
//  1. load .laser/config.yaml
//  2. parse .laser/data.yaml
//  3. iterate over all pushed files
func Handle(laser *state.L) *state.L {
	for _, l := range laser.SyncedFiles {
		processSyncedFile(laser.Scope, l, laser.StaticVars)
	}
	return laser
}

// 1. parse file
// 2. substitute vars from .laser/config.yaml on parsed file
// 3. output parsed file to path arg
func processSyncedFile(scope string, syncedFile state.SyncedFile, staticVars []state.StaticVar) {
	slog.Info("iterating over file", "path", syncedFile.Path, "id", syncedFile.ID, "hash", syncedFile.Hash)
	f, err := os.Open(path.Join(scope, files.DefautSyncedFolder, syncedFile.Hash))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	untransformed, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	subs := []string{}
	for _, v := range staticVars {
		slog.Info("variable substitution will occur", "name", v.Name, "value", v.Value)
		subs = append(subs, fmt.Sprintf("{{%s}}", v.Name), v.Value)
	}

	replacer := strings.NewReplacer(subs...)
	transformed := replacer.Replace(string(untransformed))

	if err := os.WriteFile(syncedFile.Path, []byte(transformed), os.ModePerm); err != nil {
		panic(err)
	}
}
