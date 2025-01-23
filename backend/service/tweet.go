package service

import (
	"context"

	"github.com/Takenari-Yamamoto/twitter-clone/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TweetService struct {
	db boil.ContextExecutor
}

func NewTweetService(db boil.ContextExecutor) *TweetService {
	return &TweetService{db: db}
}

func (s *TweetService) CreateTweet(ctx context.Context, userID int64, content string) (*models.Tweet, error) {
	tweet := &models.Tweet{
		UserID:  userID,
		Content: content,
	}

	if err := tweet.Insert(ctx, s.db, boil.Infer()); err != nil {
		return nil, err
	}

	// ユーザー情報を含めて再取得
	tweet, err := models.FindTweet(ctx, s.db, tweet.ID)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
