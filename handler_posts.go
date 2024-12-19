package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/romusking/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		arg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return err
		}
		if arg < 1 {
			return fmt.Errorf("limit can't be less than 1")
		}
		limit = arg
	}
	params := database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetPosts(context.Background(), params)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("---------------------------\nTitle: %s\n\nDescription: %s\nLink: %s\nPublished: %q\n\n\n", post.Title, post.Description, post.Url, post.PublishedAt)
	}
	return nil
}
