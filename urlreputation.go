package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/url_intel"
	"go.riyazali.net/sqlite"
)

type URLReputation struct{}

func (*URLReputation) Args() int { return 2 }

func (*URLReputation) Deterministic() bool { return true }

func (*URLReputation) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
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
