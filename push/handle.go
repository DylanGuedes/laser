package push

import (
	"github.com/DylanGuedes/laser/state"
	"golang.org/x/exp/slog"
)

func (p *PushCmd) Handle(laser *state.L) *state.L {
	fileID := *p.idFlag
	filePath := *p.pathFlag
	slog.Info("push command invoked", "file_id", fileID, "path", filePath)

	syncedFile := state.NewSyncedFile(fileID, filePath)
	syncedFile.MkCopy(laser.Scope)

	laser.AddSyncFile(syncedFile)
	laser.StoreState()

	return laser
}
