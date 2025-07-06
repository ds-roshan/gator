package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ds-roshan/gator/internal/config"
	"github.com/ds-roshan/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading a file:", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("Error opening SQL: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)
	st := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerDeleteAllUser)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAggrate)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	command := command{
		Name: args[1],
		Args: args[2:],
	}

	err = cmds.run(st, command)
	if err != nil {
		log.Fatal(err)
	}

}
