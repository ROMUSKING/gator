package main

import (
	"context"
	"fmt"

	"github.com/romusking/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	rss, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(rss)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	paramsFeed := database.CreateFeedParams{
		UserID: user.ID,
		Name:   cmd.Args[0],
		Url:    cmd.Args[1],
	}
	feed, err := s.db.CreateFeed(context.Background(), paramsFeed)
	if err != nil {
		return err
	}
	
	paramsFollow := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), paramsFollow)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(feeds)
	return nil

}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("expected: follow <url>")
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	params := database.CreateFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println(follow.FeedName, follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}
