package twitterscraper

import (
	"encoding/json"
	"errors"
	"fmt"
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
        var apiError struct {
            Errors []struct {
                Code    int    `json:"code"`
                Message string `json:"message"`
            } `json:"errors"`
        }
        if jsonErr := json.Unmarshal([]byte(err.Error()), &apiError); jsonErr == nil {
            for _, e := range apiError.Errors {
                switch e.Code {
                case 160: // Already following
                    return fmt.Errorf("Already following user: %s", screenName)
                case 161: // User is protected
                    return fmt.Errorf("User %s is protected", screenName)
                default:
                    return fmt.Errorf("Error following user: %s - %s", screenName, e.Message)
                }
            }
        }
        return err
    }

    return nil
}
