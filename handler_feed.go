package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/romusking/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		fmt.Println("usage: agg <time_between_reqs>, minimum 1m")
		return fmt.Errorf("usage: agg <time_between_reqs>, minimum 1m")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("wrong time format: %v", err)
	}
	time_between_reqs = max(time.Minute, time_between_reqs)
	fmt.Println("Collecting feeds every", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
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

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("expected: follow <url>")
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
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

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("expected: follow <url>")
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Println(feed.Name, "unfollowed")
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func scrapeFeeds(s *state) error {

	feed, err := s.db.GetAndMarkFeed(context.Background())
	if err != nil {
		return err
	}
	items, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	const timeFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

	for _, item := range items.Channel.Item {
		published, err := time.Parse(timeFormat, item.PubDate)
		if err != nil {
			log.Fatalf("item %s date in wrong format: %v", item.Title, err)
		}
		params := database.CreatePostParams{
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			FeedID:      feed.ID,
			PublishedAt: published.UTC(),
		}
		err = s.db.CreatePost(context.Background(), params)
		if err != nil {
			if err.Error() == `pq: duplicate key value violates unique constraint "posts_url_key"` {
				continue
			} else {
				log.Println(err)
			}
		}
	}
	return nil
}
