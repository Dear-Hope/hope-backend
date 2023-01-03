package storyroom

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/storyroom/filter"
)

type Repository interface {
	StorePost(newPost model.Post) (*model.Post, error)
	StorePostCategories(postID uint, categoryIDs []uint) (model.PostCategories, error)
	StoreComment(newComment model.Comment) (*model.Comment, error)
	GetAll(f filter.List) (model.Posts, error)
	GetByID(id, userID uint) (*model.Post, error)
	GetAllCommentByPostID(postID uint) (model.Comments, error)
	GetAllCategoryByPostID(postID uint) (model.PostCategories, error)
	GetTotalCommentsByPostID(postID uint) int
	GetTotalLikesByPostID(postID uint) int
	CheckIsPostLikedByUserID(postID, userID uint) bool
	UpsertLike(postID, userID uint) error
	DeletePostByID(postID uint) error
	DeleteCommentByPostID(postID uint) error
	DeletePostCategoryByPostID(postID uint) error
	DeleteLikeByPostID(postID uint) error
	GetAllCategories() (model.PostCategories, error)
	GetAllReason() (model.ReportReasons, error)
	StoreReport(newReport model.Report) (*model.Report, error)
}
