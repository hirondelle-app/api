package tweets

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Manager struct {
	*gorm.DB `inject:""`
}

func (m *Manager) GetTweetsByUser() ([]Tweet, error) {
	tweetList := []Tweet{}
	m.Preload("Keyword").Find(&tweetList)
	return tweetList, nil
}

func (m *Manager) CreateTweet(tweet *Tweet) error {
	if err := m.Create(&tweet).Error; err != nil {
		return err
	}
	m.Model(&tweet).Related(&Keyword{}, "KeywordID")
	m.Preload("Keyword").Model(&tweet)
	return nil
}

func (m *Manager) ValidateTweet(tweet *Tweet) error {
	if tweet.TweetID == "" {
		return errors.New("Tweet_id must not be empty")
	}
	if tweet.Likes == 0 {
		return errors.New("Likes must not be empty")
	}
	if tweet.Retweets == 0 {
		return errors.New("Retweets must not be empty")
	}
	if tweet.KeywordID == 0 {
		return errors.New("Keyword must not be empty")
	}
	return nil
}

func (m *Manager) CreateKeyword(keyword *Keyword) error {
	if err := m.Create(&keyword).Error; err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetKeywords() ([]Keyword, error) {
	keywords := []Keyword{}
	m.Preload("Tweets").Find(&keywords)
	return keywords, nil
}
