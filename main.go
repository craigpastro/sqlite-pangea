package main

import (
	"go.riyazali.net/sqlite"
)

// Version is set in Makefile
var Version string

type VersionFunc struct{}

func (*VersionFunc) Deterministic() bool { return true }
func (*VersionFunc) Args() int           { return 0 }
func (*VersionFunc) Apply(c *sqlite.Context, _ ...sqlite.Value) {
	c.ResultText(Version)
}

func init() {
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("pangea_version", &VersionFunc{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("redact", &Redact{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("ip_intel", &IPIntel{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("url_reputation", &URLReputation{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}

func main() {}
