package main

import "fmt"

type command struct {
	Name string
	Args []string
}
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	err := c.registeredCommands[cmd.Name](s, cmd)
	if err != nil {
		return fmt.Errorf("can't execute %s: %v", cmd.Name, err)
	}
	return nil
}
