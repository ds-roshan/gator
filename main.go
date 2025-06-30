package main

import (
	"fmt"
	"log"

	"github.com/ds-roshan/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading a file:", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("hihi")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("config: %+v\n", cfg)

}
