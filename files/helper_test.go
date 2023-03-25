package files_test

import (
	"os"
	"path"
	"testing"

	"github.com/DylanGuedes/laser/files"
	"github.com/stretchr/testify/require"
)

func TestFiles_Available(t *testing.T) {
	// check that initially it isn't available.
	scope := "fixture"
	require.False(t, files.Available(scope))
	require.NoError(t, os.Mkdir(scope, os.ModeTemporary|os.ModePerm))
	_, err := os.ReadDir(scope)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(scope))
	})

	// create configuration file.
	p := path.Join(scope, files.Config)
	f, err := os.Create(p)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, f.Close())
	})

	// check that it is available now.
	require.True(t, files.Available(scope))
}
