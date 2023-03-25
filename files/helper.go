package files

import (
	"os"
	"path"

	"golang.org/x/exp/slog"
)

func Available(scope string) bool {
	info, err := os.Stat(path.Join(scope, Config))
	if err != nil {
		slog.Error("os stat failed", "scope", scope, "file", Config)
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}
