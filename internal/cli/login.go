package cli

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("login requires exactly one username")
	}

	username := cmd.Args[0]

	if err := s.Config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", username)
	return nil
}
