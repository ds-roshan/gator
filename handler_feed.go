package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {

	currentUser := s.cfg.CurrentUserName
	usr, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return errors.New("invalid user: login again or register a user")
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage %s <title> <url>", cmd.Name)
	}

	title := cmd.Args[0]
	url := cmd.Args[1]

	feedParam := database.CreateFeedParams{
		ID: uuid.New(),
		Name: sql.NullString{
			String: title,
			Valid:  title != "",
		},
		Url: sql.NullString{
			String: url,
			Valid:  url != "",
		},
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: uuid.NullUUID{
			UUID:  usr.ID,
			Valid: true,
		},
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParam)
	if err != nil {
		return fmt.Errorf("couldn't create feed %w", err)
	}

	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("================================")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:			%s\n", feed.ID)
	fmt.Printf("* Created:		%v\n", feed.CreatedAt)
	fmt.Printf("* Updated:		%v\n", feed.UpdatedAt)
	fmt.Printf("* Name:			%s\n", feed.Name)
	fmt.Printf("* URL:			%s\n", feed.Url)
	fmt.Printf("* UserID:		%s\n", feed.UserID)
}
