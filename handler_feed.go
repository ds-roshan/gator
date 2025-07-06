package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <title> <url>", cmd.Name)
	}

	title := cmd.Args[0]
	url := cmd.Args[1]

	feedParam := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      title,
		Url:       url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return fmt.Errorf("couldn't create feed %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed, user)
	fmt.Println()
	fmt.Println("================================")

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("You're following %s\n", feedFollow.FeedName)

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		usr, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeed(feed, usr)
		println("================================")
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* Created:		%v\n", feed.CreatedAt)
	fmt.Printf("* Updated:		%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
	fmt.Printf("* User:			%s\n", user.Name)
}
