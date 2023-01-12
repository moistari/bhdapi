// example/example.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/moistari/bhdapi"
)

func main() {
	api := flag.String("api", "", "api key")
	rss := flag.String("rss", "", "rss key")
	flag.Parse()
	cl := bhdapi.New(
		bhdapi.WithApiKey(*api),
		bhdapi.WithRssKey(*rss, false),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res, err := cl.Search(ctx, "fight club framestor")
	if err != nil {
		log.Fatal(err)
	}
	var torrent bhdapi.Torrent
	found := false
	for count, page, hasMore := 0, 1, true; hasMore; page++ {
		for _, r := range res.Results {
			if r.Name == "Fight Club 1999 BluRay 1080p DTS-HD MA 5.1 AVC REMUX-FraMeSToR" {
				found, torrent = true, r
				break
			}
			count++
		}
		if found {
			break
		}
		hasMore = count < res.TotalResults
	}
	if !found {
		log.Fatal("could not find torrent")
	}
	fmt.Printf("torrent: %d: %s (%d)\n", torrent.ID, torrent.Name, torrent.Size)
	buf, err := cl.Torrent(ctx, torrent.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved torrent with length: %d\n", len(buf))
}
