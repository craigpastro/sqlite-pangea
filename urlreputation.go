package main

import (
	"context"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/url_intel"
	"go.riyazali.net/sqlite"
	"gopkg.in/square/go-jose.v2/json"
)

type URLReputation struct{}

func (r *URLReputation) Args() int { return 2 }

func (r *URLReputation) Deterministic() bool { return true }

func (r *URLReputation) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	token := values[0].Text()
	url := values[1].Text()

	client := url_intel.New(&pangea.Config{
		Token:  token,
		Domain: "aws.us.pangea.cloud",
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
