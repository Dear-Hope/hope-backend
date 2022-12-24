package model

import (
	"time"
)

type (
	Post struct {
		ID            uint      `db:"id"`
		AuthorID      uint      `db:"author_id"`
		AuthorName    string    `db:"author_name"`
		AuthorPicURL  string    `db:"author_pic_url"`
		Content       string    `db:"content"`
		TotalComments int       `db:"total_comments"`
		TotalLikes    int       `db:"total_likes"`
		Liked         bool      `db:"liked"`
		CreatedAt     time.Time `db:"created_at"`

		Comments   Comments
		Categories PostCategories
	}

	Posts []Post

	Comment struct {
		ID           uint      `db:"id"`
		PostID       uint      `db:"post_id"`
		AuthorID     uint      `db:"author_id"`
		AuthorName   string    `db:"author_name"`
		AuthorPicURL string    `db:"author_pic_url"`
		Content      string    `db:"content"`
		CreatedAt    time.Time `db:"created_at"`
	}

	Comments []Comment

	PostCategory struct {
		ID       uint   `db:"id"`
		Name     string `db:"name"`
		ImageURL string `db:"image_url"`
	}

	PostCategories []PostCategory
)

func (ths Post) TableWithSchemaName() string {
	return `"storyroom".posts`
}

func (ths Post) ToPostResponse() *PostResponse {
	var categories []string

	if ths.Categories != nil {
		for _, category := range ths.Categories {
			categories = append(categories, category.Name)
		}
	}

	return &PostResponse{
		ID: ths.ID,
		Author: Author{
			ID:        ths.AuthorID,
			Name:      ths.AuthorName,
			AvatarURL: ths.AuthorPicURL,
		},
		Content:       ths.Content,
		TotalComments: ths.TotalComments,
		TotalLikes:    ths.TotalLikes,
		Liked:         ths.Liked,
		CreatedAt:     ths.CreatedAt.UTC(),
		Comments:      ths.Comments.ToListCommentResponse(),
		Categories:    categories,
	}
}

func (ths Posts) ToListPostResponse() []PostResponse {
	res := make([]PostResponse, len(ths))
	for i, post := range ths {
		res[i] = *post.ToPostResponse()
	}

	return res
}

func (ths Comments) ToListCommentResponse() []CommentResponse {
	res := make([]CommentResponse, len(ths))
	for i, comment := range ths {
		res[i] = *comment.ToCommentResponse()
	}

	return res
}

func (ths Comment) ToCommentResponse() *CommentResponse {
	return &CommentResponse{
		ID: ths.ID,
		Author: Author{
			ID:        ths.AuthorID,
			Name:      ths.AuthorName,
			AvatarURL: ths.AuthorPicURL,
		},
		Content:   ths.Content,
		CreatedAt: ths.CreatedAt.UTC(),
	}
}

func (ths PostCategory) ToPostCategoryResponse() *PostCategoryResponse {
	return &PostCategoryResponse{
		ID:       ths.ID,
		Name:     ths.Name,
		ImageURL: ths.ImageURL,
	}
}

func (ths PostCategories) ToListPostCategoryResponse() []PostCategoryResponse {
	res := make([]PostCategoryResponse, len(ths))
	for i, postCategory := range ths {
		res[i] = *postCategory.ToPostCategoryResponse()
	}

	return res
}

type (
	Author struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		AvatarURL string `json:"avatarUrl"`
	}

	PostResponse struct {
		ID            uint              `json:"id"`
		Author        Author            `json:"author"`
		Content       string            `json:"content"`
		TotalComments int               `json:"totalComments"`
		TotalLikes    int               `json:"totalLikes"`
		Liked         bool              `json:"liked"`
		CreatedAt     time.Time         `json:"createdAt"`
		Comments      []CommentResponse `json:"comments"`
		Categories    []string          `json:"categories"`
	}

	CommentResponse struct {
		ID        uint      `json:"id"`
		Author    Author    `json:"author"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"createdAt"`
	}

	PostCategoryResponse struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		ImageURL string `json:"imageUrl"`
	}

	PostRequest struct {
		Content     string `json:"content"`
		CategoryIDs []uint `json:"categoryIds"`
	}

	CommentRequest struct {
		Content string `json:"content"`
	}
)