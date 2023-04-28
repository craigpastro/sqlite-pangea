package main

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"crawshaw.io/sqlite"
	"crawshaw.io/sqlite/sqlitex"
	"github.com/stretchr/testify/require"
)

const extension = "pangea.so"

func TestRedact(t *testing.T) {
	token, ok := os.LookupEnv("PANGEA_TOKEN")
	if !ok {
		t.Fatal("please set the PANGEA_TOKEN environment variable")
	}

	if _, err := os.Stat(extension); errors.Is(err, os.ErrNotExist) {
		t.Fatal("'pangea.so' does not exist; perhaps you need to build it?")
	}

	conn, err := sqlite.OpenConn("file::memory:", 0)
	require.NoError(t, err)

	err = conn.EnableLoadExtension(true)
	require.NoError(t, err)

	err = conn.LoadExtension(extension, "")
	require.NoError(t, err)

	q := fmt.Sprintf("select redact('%s', 'my phone number is 123-456-7890')", token)

	stmt, err := conn.Prepare(q)
	require.NoError(t, err)

	got, err := sqlitex.ResultText(stmt)
	require.NoError(t, err)
	require.Equal(t, "my phone number is <PHONE_NUMBER>", got)
}
