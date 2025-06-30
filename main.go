package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ds-roshan/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading a file:", err)
	}

	st := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

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
