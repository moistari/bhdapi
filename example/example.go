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
	req := bhdapi.Search("fight club framestor")
	var found *bhdapi.Torrent
	for req.Next(ctx, cl) {
		torrent := req.Cur()
		if torrent.Name == "Fight Club 1999 BluRay 1080p DTS-HD MA 5.1 AVC REMUX-FraMeSToR" {
			found = &torrent
			break
		}
	}
	if err := req.Err(); err != nil {
		log.Fatal(err)
	}
	if found == nil {
		log.Fatal("could not find torrent")
	}
	fmt.Printf("found: %d: %s (%d)\n", found.ID, found.Name, found.Size)
	buf, err := cl.Torrent(ctx, found.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved found torrent with length: %d\n", len(buf))
}
