package handler

import (
	"fmt"

	"github.com/Takenari-Yamamoto/twitter-clone/gen/models"
	"github.com/Takenari-Yamamoto/twitter-clone/gen/restapi/operations/tweets"
	"github.com/Takenari-Yamamoto/twitter-clone/service"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type TweetHandler struct {
	tweetService *service.TweetService
}

func NewTweetHandler(tweetService *service.TweetService) *TweetHandler {
	if tweetService == nil {
		panic("tweetService must not be nil")
	}
	return &TweetHandler{
		tweetService: tweetService,
	}
}

func (h *TweetHandler) HandleCreateTweet(params tweets.PostTweetsParams, principal interface{}) middleware.Responder {
	fmt.Println("Creating tweet...")
	// TODO: principalからユーザーIDを取得する処理を実装
	userID := int64(1) // 仮の実装

	tweet, err := h.tweetService.CreateTweet(params.HTTPRequest.Context(), userID, *params.Body.Content)
	if err != nil {
		return tweets.NewPostTweetsInternalServerError().WithPayload(&models.Error{
			Code:    500,
			Message: err.Error(),
		})
	}

	return tweets.NewPostTweetsCreated().WithPayload(&models.Tweet{
		ID:        tweet.ID,
		Content:   tweet.Content,
		UserID:    tweet.UserID,
		CreatedAt: strfmt.DateTime(tweet.CreatedAt),
		UpdatedAt: strfmt.DateTime(tweet.UpdatedAt),
	})
}
