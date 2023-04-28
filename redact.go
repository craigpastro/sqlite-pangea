package main

import (
	"context"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/redact"
	"go.riyazali.net/sqlite"
)

type Redact struct{}

func (r *Redact) Args() int { return 2 }

func (r *Redact) Deterministic() bool { return true }

func (r *Redact) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	token := values[0].Text()
	text := values[1].Text()

	client := redact.New(&pangea.Config{
		Token:  token,
		Domain: "aws.us.pangea.cloud",
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
