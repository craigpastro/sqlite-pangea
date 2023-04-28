package main

import (
	"go.riyazali.net/sqlite"
)

func init() {
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		if err := api.CreateFunction("redact", &Redact{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("url_reputation", &URLReputation{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}

func main() {}
