package main

import (
	"context"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/redact"
	"go.riyazali.net/sqlite"
)

type Redact struct {
	conn *sqlite.Conn
}

func (r *Redact) GetConn() *sqlite.Conn {
	return r.conn
}

func (*Redact) Args() int { return 1 }

func (*Redact) Deterministic() bool { return true }

func (r *Redact) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	domain, token, err := getPangeaDomainAndToken(r.conn)
	if err != nil {
		ctx.ResultError(fmt.Errorf("failed to retrieve the config: %w", err))
		return
	}

	text := values[0].Text()

	client := redact.New(&pangea.Config{
		Domain: domain,
		Token:  token,
	})

	resp, err := client.Redact(context.Background(), &redact.TextInput{
		Text: pangea.String(text),
		// Rules: []string{},
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea error: %w", err))
		return
	}

	ctx.ResultText(*resp.Result.RedactedText)
}
