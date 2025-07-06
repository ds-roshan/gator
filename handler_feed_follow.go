package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usages %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

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

func handlerFollowing(s *state, _ command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)

	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf("ID: %s\n", feed.ID)
		fmt.Printf("Feed Name: %s\n", feed.FeedName)
		fmt.Printf("User: %s\n", feed.UserName)
		fmt.Println("====================")
	}
	return nil
}
