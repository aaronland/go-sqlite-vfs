package vfs

import (
	"context"
	"fmt"
	"github.com/psanford/sqlite3vfs"
	"github.com/psanford/sqlite3vfshttp"
	"strings"
)

func NewDSN(ctx context.Context, dsn string) (string, error) {

	if !strings.HasPrefix(dsn, "http") {
		return "", fmt.Errorf("DSN does not start with `http`")
	}

	vfs := sqlite3vfshttp.HttpVFS{
		URL: dsn,
		// RoundTripper: &roundTripper{}
	}

	err := sqlite3vfs.RegisterVFS("httpvfs", &vfs)

	if err != nil {
		return "", fmt.Errorf("Failed to register VFS, %w", err)
	}

	// See the vfs: prefix? We check that in sqlite.go to determine
	// whether or not to os.Stat the database

	dsn = "vfs:stub.db?vfs=httpvfs&mode=ro"

	return dsn, nil
}
