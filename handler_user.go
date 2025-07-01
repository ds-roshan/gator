package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ds-roshan/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("username is required")
	}

	username := cmd.Args[0]

	usr, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("No user found: %v", username)
	}

	err = s.cfg.SetUser(usr.Name)
	if err != nil {
		return fmt.Errorf("Couldn't set current user: %w", err)
	}

	fmt.Printf("User switched to %v\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {

	if len(cmd.Args) != 1 {
		return errors.New("Type your username to register")
	}

	name := cmd.Args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	usr, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	s.cfg.SetUser(usr.Name)
	fmt.Printf("User was created: %v\n", usr.Name)

	return nil
}

func handlerDeleteAllUser(s *state, cmd command) error {
	return s.db.DeleteAllUser(context.Background())
}

func handlerGetUsers(s *state, cmd command) error {
	usrs, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	currentUsr := s.cfg.CurrentUserName

	for _, usr := range usrs {
		if usr.Name == currentUsr {
			fmt.Printf("* %v (current)\n", usr.Name)
			continue
		}
		fmt.Printf("* %v\n", usr.Name)
	}

	return nil
}
