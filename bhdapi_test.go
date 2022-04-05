package bhdapi

import (
	"bytes"
	"context"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	cl := buildClient()
	res, err := cl.Search(context.Background(), "fight club remux framestor")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	for i, r := range res.Results {
		t.Logf("%02d: %s %d %q", i, r.InfoHash, r.ID, r.Name)
	}
}

func TestTorrent(t *testing.T) {
	cl := buildClient()
	res, err := cl.Torrent(context.Background(), 7531)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if !bytes.Contains(res, []byte("Fight.Club.1999.BluRay.1080p.DTS-HD.MA.5.1.AVC.REMUX-FraMeSToR")) {
		t.Errorf("expected buf to contain torrent name")
	}
}

func buildClient() *Client {
	var opts []Option
	if apiKey := os.Getenv("APIKEY"); apiKey != "" {
		opts = append(opts, WithApiKey(apiKey))
	}
	if rssKey := os.Getenv("RSSKEY"); rssKey != "" {
		opts = append(opts, WithRssKey(rssKey, false))
	}
	return New(opts...)
}
