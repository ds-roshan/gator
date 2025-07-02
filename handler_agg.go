package main

import (
	"context"
	"fmt"
)

func handlerAggrate(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("RSSFeed: %+v\n", feed)
	return nil
}
