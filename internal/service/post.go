package service

import (
	"voice-out-be/internal/dto"
	"voice-out-be/internal/model"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/storage"

	"go.uber.org/zap"
)

type postService struct {
	config  Config
	storage storage.Storage
}

func NewPostService(config Config, store storage.Storage) *postService {

	return &postService{
		config:  config,
		storage: store,
	}
}

func (s *postService) CreatePost(createPostReq *request.CreatePostRequest, userID string) (code response.Code, err error) {

	s.config.Logger().Info("CreatePost: creating post")

	postIn := &dto.CreatePost{
		From:   createPostReq.From,
		To:     createPostReq.To,
		Body:   createPostReq.Body,
		UserID: userID,
	}

	_, err = s.storage.CreatePost(postIn)

	if err != nil {
		s.config.Logger().Error("CreatePost: error creating post", zap.Error(err))
		return response.ServerError, err
	}

	return response.Success, nil
}

func (s *postService) FetchAllPosts() (code response.Code, posts *[]model.Post, err error) {

	s.config.Logger().Info("FetchPosts: fetching all posts")

	posts, err = s.storage.FindAllPosts()

	if err != nil {
		s.config.Logger().Error("FetchPosts: error fetching all posts", zap.Error(err))
		return response.ServerError, nil, err
	}

	return response.Success, posts, nil
}
