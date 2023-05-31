package main

import (
	"errors"

	"go.riyazali.net/sqlite"
)

func getPangeaDomainAndToken(conn *sqlite.Conn) (domain, token string, err error) {
	stmt, _, err := conn.Prepare("select domain, token from pangea_config order by id desc limit 1;")
	if err != nil {
		return "", "", err
	}

	ok, err := stmt.Step()
	if !ok || err != nil {
		return "", "", errors.New("cannot retrieve values from pange_config table")

	}

	domain = stmt.GetText("domain")
	token = stmt.GetText("token")
	_ = stmt.Finalize()

	return
}

func init() {
	sqlite.Register(func(api *sqlite.ExtensionApi) (sqlite.ErrorCode, error) {
		conn := api.Connection()
		err := conn.Exec(`create table pangea_config (
			id integer primary key,
			domain text,
			token text
		);`, func(stmt *sqlite.Stmt) error {
			return nil
		})
		if err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("pangea_version", &VersionFunc{}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("redact", &Redact{conn: conn}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("ip_intel", &IPIntel{conn: conn}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		if err := api.CreateFunction("url_reputation", &URLReputation{conn: conn}); err != nil {
			return sqlite.SQLITE_ERROR, err
		}

		return sqlite.SQLITE_OK, nil
	})
}

func main() {}
