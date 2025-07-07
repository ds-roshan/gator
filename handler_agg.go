package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAggrate(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("usages %s <req_gap 5s, 5m,>", cmd.Name)
	}

	time_between_reqs := cmd.Args[0]
	t, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(t)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feed", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Println("couldn't mark feed as fetched", err)
		return
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Println("couldn't fetch the feed", err)
		return
	}

	for _, feedItem := range feed.Channel.Item {
		updateToDB(s, feedItem, nextFeed)
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(feed.Channel.Item))
}

func updateToDB(s *state, item RSSItem, feed database.Feed) {

	t, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		panic(err)
	}

	_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title: sql.NullString{
			String: item.Title,
			Valid:  item.Title != "",
		},
		Url: item.Link,
		Description: sql.NullString{
			String: item.Description,
		},
		PublishedAt: sql.NullTime{
			Time:  t,
			Valid: true,
		},
		FeedID: feed.ID,
	})

	if err != nil {
		log.Printf("Couldn't create a post %v", err)
	}

	fmt.Println("Post created successfully")

}
