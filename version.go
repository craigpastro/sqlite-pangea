package main

import "go.riyazali.net/sqlite"

// Version is set in Makefile
var Version string

type VersionFunc struct{}

func (*VersionFunc) Deterministic() bool { return true }
func (*VersionFunc) Args() int           { return 0 }
func (*VersionFunc) Apply(c *sqlite.Context, _ ...sqlite.Value) {
	c.ResultText(Version)
}
