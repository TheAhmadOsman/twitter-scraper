package twitterscraper_test

import (
	"strings"
	"testing"
	"time"

	twitterscraper "github.com/TheAhmadOsman/twitter-scraper"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestGetProfile(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	joined := time.Date(2010, 01, 18, 8, 49, 30, 0, loc)
	sample := twitterscraper.Profile{
		Avatar:    "https://pbs.twimg.com/profile_images/436075027193004032/XlDa2oaz_normal.jpeg",
		Banner:    "https://pbs.twimg.com/profile_banners/106037940/1541084318",
		Biography: "nothing",
		//	Birthday:   "March 21",
		IsPrivate:      false,
		IsVerified:     false,
		Joined:         &joined,
		Location:       "Ukraine",
		Name:           "Nomadic",
		PinnedTweetIDs: []string{},
		URL:            "https://twitter.com/nomadic_ua",
		UserID:         "106037940",
		Username:       "nomadic_ua",
		Website:        "https://nomadic.name",
	}

	profile, err := testScraper.GetProfile("nomadic_ua")
	if err != nil {
		t.Error(err)
	}

	cmpOptions := cmp.Options{
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FollowersCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FollowingCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FriendsCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "LikesCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "ListedCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "TweetsCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "MediaCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "NormalFollowersCount"),
	}
	if diff := cmp.Diff(sample, profile, cmpOptions...); diff != "" {
		t.Error("Resulting profile does not match the sample", diff)
	}

	if profile.FollowersCount == 0 {
		t.Error("Expected FollowersCount is greater than zero")
	}
	if profile.FollowingCount == 0 {
		t.Error("Expected FollowingCount is greater than zero")
	}
	if profile.LikesCount == 0 {
		t.Error("Expected LikesCount is greater than zero")
	}
	if profile.TweetsCount == 0 {
		t.Error("Expected TweetsCount is greater than zero")
	}
}

func TestGetProfilePrivate(t *testing.T) {
	loc := time.FixedZone("UTC", 0)
	joined := time.Date(2020, 1, 26, 0, 3, 5, 0, loc)
	sample := twitterscraper.Profile{
		Avatar:    "https://pbs.twimg.com/profile_images/1612213936082030594/_HEsjv7Q_normal.jpg",
		Banner:    "https://pbs.twimg.com/profile_banners/1221221876849995777/1673110776",
		Biography: "t h e h e r m i t",
		//	Birthday:   "March 21",
		IsPrivate:      true,
		IsVerified:     false,
		Joined:         &joined,
		Location:       "sometimes",
		Name:           "private account",
		PinnedTweetIDs: []string{},
		URL:            "https://twitter.com/tomdumont",
		UserID:         "1221221876849995777",
		Username:       "tomdumont",
		Website:        "",
	}

	// some random private profile (found via google)
	profile, err := testScraper.GetProfile("tomdumont")
	if err != nil {
		t.Error(err)
	}

	cmpOptions := cmp.Options{
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FollowersCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FollowingCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "FriendsCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "LikesCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "ListedCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "TweetsCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "MediaCount"),
		cmpopts.IgnoreFields(twitterscraper.Profile{}, "NormalFollowersCount"),
	}
	if diff := cmp.Diff(sample, profile, cmpOptions...); diff != "" {
		t.Error("Resulting profile does not match the sample", diff)
	}

	if profile.FollowingCount == 0 {
		t.Error("Expected FollowingCount is greater than zero")
	}
	if profile.LikesCount == 0 {
		t.Error("Expected LikesCount is greater than zero")
	}
	if profile.TweetsCount == 0 {
		t.Error("Expected TweetsCount is greater than zero")
	}
}

func TestGetProfileErrorSuspended(t *testing.T) {
	_, err := testScraper.GetProfile("1")
	if err == nil {
		t.Error("Expected Error, got success")
	} else {
		if !strings.Contains(err.Error(), "suspended") {
			t.Error("Expected error to contain 'suspended', got", err)
		}
	}
}

func TestGetProfileErrorNotFound(t *testing.T) {
	neUser := "sample3123131"
	expectedError := "user not found"
	_, err := testScraper.GetProfile(neUser)
	if err == nil {
		t.Error("Expected Error, got success")
	} else {
		if err.Error() != expectedError {
			t.Errorf("Expected error '%s', got '%s'", expectedError, err)
		}
	}
}

func TestGetProfileByID(t *testing.T) {
	profile, err := testScraper.GetProfileByID("1221221876849995777")
	if err != nil {
		t.Error(err)
	}

	if profile.Username != "tomdumont" {
		t.Errorf("Expected username 'tomdumont', got '%s'", profile.Username)
	}
}

func TestGetUserIDByScreenName(t *testing.T) {
	userID, err := testScraper.GetUserIDByScreenName("X")
	if err != nil {
		t.Errorf("getUserByScreenName() error = %v", err)
	}
	if userID == "" {
		t.Error("Expected non-empty user ID")
	}
}
