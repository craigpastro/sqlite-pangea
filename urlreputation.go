package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/url_intel"
	"go.riyazali.net/sqlite"
)

type URLReputation struct {
	conn *sqlite.Conn
}

func (*URLReputation) Args() int { return 1 }

func (*URLReputation) Deterministic() bool { return true }

func (u *URLReputation) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	domain, token, err := getPangeaDomainAndToken(u.conn)
	if err != nil {
		ctx.ResultError(fmt.Errorf("failed to retrieve the config: %w", err))
		return
	}

	url := values[0].Text()

	client := url_intel.New(&pangea.Config{
		Domain: domain,
		Token:  token,
	})

	resp, err := client.Reputation(context.Background(), &url_intel.UrlReputationRequest{
		Url:      url,
		Provider: "crowdstrike",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea error: %w", err))
		return
	}

	b, err := json.Marshal(resp.Result.Data)
	if err != nil {
		ctx.ResultError(fmt.Errorf("json marshal: %w", err))
		return
	}

	ctx.ResultText(string(b))
}
