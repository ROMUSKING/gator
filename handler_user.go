package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if cmd.Name != "login" {
		return fmt.Errorf("program error, not login")
	}
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username required to log in")
	}
	dbUser, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatalf("user doesn't exist: %v", err)
	}
	err = s.config.SetUser(dbUser.Name)
	if err != nil {
		return fmt.Errorf("can't log in: %v", err)
	}
	fmt.Printf("User %s logged on\n", s.config.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if cmd.Name != "register" {
		return fmt.Errorf("program error, not register")
	}
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username required to register")
	}
	dbUser, err := s.db.CreateUser(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatalf("error creating user: %v", err)
	}
	s.config.SetUser(dbUser.Name)
	fmt.Println("User created successfully")
	log.Printf("User %s created", s.config.CurrentUserName)

	return nil
}

func handlerReset(s *state, cmd command) error {
	if cmd.Name != "reset" {
		return fmt.Errorf("program error, not reset")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatalf("error deleting users: %v", err)
	}

	fmt.Println("Users deleted successfully")
	log.Printf("Users deleted")

	return nil
}

func handlerUsers(s *state, cmd command) error {
	if cmd.Name != "users" {
		return fmt.Errorf("program error, not users")
	}
	dbUsers, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatalf("can't get users from db: %v", err)
	}
	currUser := s.config.CurrentUserName
	for _, user := range dbUsers {
		if user.Name == currUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
