package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/moistari/bhdapi"
)

func main() {
	apikey := flag.String("apikey", "", "api key")
	rsskey := flag.String("rsskey", "", "rss key")
	flag.Parse()
	if err := run(context.Background(), *apikey, *rsskey, flag.Args()...); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, apikey, rsskey string, args ...string) error {
	cl := bhdapi.New(bhdapi.WithApiKey(apikey), bhdapi.WithRssKey(rsskey))
	res, err := cl.Search(ctx, args...)
	if err != nil {
		return err
	}
	sort.Slice(res.Results, func(i, j int) bool {
		return res.Results[i].ID < res.Results[j].ID
	})
	for _, r := range res.Results {
		fmt.Fprintf(os.Stdout, "%02d: %s %q %s\n", r.ID, r.InfoHash[:7], r.Name, "https://beyond-hd.me/torrents/a."+strconv.Itoa(r.ID))
	}
	return nil
}
