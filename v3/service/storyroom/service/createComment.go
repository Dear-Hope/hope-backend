package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (ths *service) CreateComment(req model.CommentRequest, postID, authorID uint) (*model.CommentResponse, *model.ServiceError) {
	if postID <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_CREATE_COMMENT_FAILED, postID),
		}
	}

	newComment := model.Comment{
		Content:  req.Content,
		AuthorID: authorID,
		PostID:   postID,
	}

	comment, err := ths.repo.StoreComment(newComment)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_COMMENT_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf(constant.ERROR_CREATE_COMMENT_FAILED, postID),
			}
		}
	}

	return comment.ToCommentResponse(), nil
}
