package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pangeacyber/pangea-go/pangea-sdk/pangea"
	"github.com/pangeacyber/pangea-go/pangea-sdk/service/ip_intel"
	"go.riyazali.net/sqlite"
)

type IPIntelResponse struct {
	ReputationData *ip_intel.ReputationData `json:"reputationData"`
	Domain         string                   `json:"domain"`
	IsProxy        bool                     `json:"isProxy"`
	IsVPN          bool                     `json:"isVPN"`
	GeolocateData  *ip_intel.GeolocateData  `json:"geolocateData"`
}

type IPIntel struct {
	conn *sqlite.Conn
}

func (*IPIntel) Args() int { return 1 }

func (*IPIntel) Deterministic() bool { return true }

func (i *IPIntel) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	domain, token, err := getPangeaDomainAndToken(i.conn)
	if err != nil {
		ctx.ResultError(fmt.Errorf("failed to retrieve the config: %w", err))
		return
	}

	ip := values[0].Text()

	client := ip_intel.New(&pangea.Config{
		Domain: domain,
		Token:  token,
	})

	c := context.Background()

	reputation, err := client.Reputation(c, &ip_intel.IpReputationRequest{
		Ip:       ip,
		Provider: "crowdstrike",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea ip reputation error: %w", err))
		return
	}

	ipDomain, err := client.GetDomain(c, &ip_intel.IpDomainRequest{
		Ip:       ip,
		Provider: "digitalelement",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea ip get domain error: %w", err))
		return
	}

	proxy, err := client.IsProxy(c, &ip_intel.IpProxyRequest{
		Ip:       ip,
		Provider: "digitalelement",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea ip is proxy error: %w", err))
		return
	}

	vpn, err := client.IsVPN(c, &ip_intel.IpVPNRequest{
		Ip:       ip,
		Provider: "digitalelement",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea ip is vpn error: %w", err))
		return
	}

	geolocate, err := client.Geolocate(c, &ip_intel.IpGeolocateRequest{
		Ip:       ip,
		Provider: "digitalelement",
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("pangea ip geolocate error: %w", err))
		return
	}

	b, err := json.Marshal(&IPIntelResponse{
		ReputationData: &reputation.Result.Data,
		Domain:         ipDomain.Result.Data.Domain,
		IsProxy:        proxy.Result.Data.IsProxy,
		IsVPN:          vpn.Result.Data.IsVPN,
		GeolocateData:  &geolocate.Result.Data,
	})
	if err != nil {
		ctx.ResultError(fmt.Errorf("json marshal: %w", err))
		return
	}

	ctx.ResultText(string(b))
}
