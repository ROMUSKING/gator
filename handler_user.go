package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if cmd.Name != "login" {
		return fmt.Errorf("program error, not login")
	}
	if len(cmd.Args) != 1 {
		return fmt.Errorf("username required to log in")
	}
	err := s.config.SetUser(cmd.Args[0])
	if err != nil {

		return fmt.Errorf("can't log in: %v", err)
	}
	fmt.Printf("User %s logged on\n", s.config.CurrentUserName)
	return nil
}
