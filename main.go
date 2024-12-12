package main

import (
	"log"
	"os"

	"github.com/romusking/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	s := state{config: &cfg}
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("not enough arguments\n")
	}
	cmd := command{
		Name: args[1],
		Args: args[2:],
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("can't execute command: %v\n", err)
	}
}
