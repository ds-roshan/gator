package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ds-roshan/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if limitInput, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = limitInput
		} else {
			return fmt.Errorf("Limit number is not valid")
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s from %s \n", post.UpdatedAt.Format("Mon Jul 7"), post.FeedName)
		fmt.Printf("Title:			%v\n", post.Title)
		fmt.Printf("Description:	%v\n", post.Description.String)
		fmt.Printf("Link:			%v\n", post.Url)
	}
	return nil
}
