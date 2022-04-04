package bhdapi

import (
	"context"
	"os"
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
	var opts []Option
	if apiKey := os.Getenv("APIKEY"); apiKey != "" {
		opts = append(opts, WithApiKey(apiKey))
	}
	if rssKey := os.Getenv("RSSKEY"); rssKey != "" {
		opts = append(opts, WithRssKey(rssKey))
	}
	req := Search().
		WithFreeleech(true)
	res, err := req.Do(context.Background(), New(opts...))
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for i, r := range res.Results {
		t.Logf("%02d: %s %s -- %s", i, r.InfoHash, r.Name, "https://beyond-hd.me/torrents/a."+strconv.Itoa(r.ID))
	}
}
