package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/romusking/gator/internal/config"
	"github.com/romusking/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("no connection to database: %v", err)
	}
	s := state{
		config: &cfg,
		db:     database.New(db),
	}
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
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("mv", handlerMV)
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("can't execute command: %v\n", err)
	}
}
