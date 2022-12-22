package storyroom

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
)

type Service interface {
	CreatePost(req model.PostRequest, authorID uint) (*model.PostResponse, *model.ServiceError)
	CreateComment(req model.CommentRequest, postID, authorID uint) (*model.CommentResponse, *model.ServiceError)
	List(f filter.List, userID uint) ([]model.PostResponse, *model.ServiceError)
	GetByID(id, userID uint) (*model.PostResponse, *model.ServiceError)
	UpsertLike(postID, userID uint) *model.ServiceError
	DeletePostByID(postID, userID uint) *model.ServiceError
	ListCategories() ([]model.PostCategoryResponse, *model.ServiceError)
}
