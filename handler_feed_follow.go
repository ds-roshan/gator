package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usages %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)

	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    usr.ID,
	})

	fmt.Printf("You're following %s\n", feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, _ command) error {
	usr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), usr.ID)

	for _, feed := range feeds {
		fmt.Printf("ID: %s\n", feed.ID)
		fmt.Printf("Feed Name: %s\n", feed.FeedName)
		fmt.Printf("User: %s\n", feed.UserName)
		fmt.Println("====================")
	}
	return nil
}
