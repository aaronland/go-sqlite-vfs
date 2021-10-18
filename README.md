# go-sqlite-vfs

Go package to register `psanford/sqlite3vfs` Virtual File Systems and construct corresponding DSN strings.

## Motivation

This code used to be a part of the [aaronland/go-sqlite](https://github.com/aaronland/go-sqlite/) package but was removed because the [psanford/sqlite3vfs](https://github.com/psanford/sqlite3vfs) still has compile-time issues on certain platforms.

## See also

* https://github.com/aaronland/go-sqlite/
* https://github.com/psanford/sqlite3vfs