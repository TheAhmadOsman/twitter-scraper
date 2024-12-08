package twitterscraper

import (
	"errors"
	"net/url"
)

// FollowUser follows a user by their screen name.
func (s *Scraper) FollowUser(screenName string) error {
    if !s.isLogged {
        return errors.New("scraper is not logged in")
    }

    req, err := s.newRequest("POST", "https://api.twitter.com/1.1/friendships/create.json")
    if err != nil {
        return err
    }

    query := url.Values{}
    query.Set("screen_name", screenName)
    query.Set("follow", "true")
    req.URL.RawQuery = query.Encode()

    var response struct{}
    err = s.RequestAPI(req, &response)
    if err != nil {
        return err
    }

    return nil
}
