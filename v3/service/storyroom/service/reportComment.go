package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (ths *service) ReportComment(req model.ReportCommentRequest, userID uint) (*model.ReportCommentResponse, *model.ServiceError) {
	if req.CommentID <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_REPORT_FAILED, "comment", req.CommentID),
		}
	}

	if req.ReasonID <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_REPORT_POST_MINIMAL_ONE_REASON),
		}
	}

	newReport := model.ReportComment{
		UserID:    userID,
		CommentID: req.CommentID,
		ReasonID:  req.ReasonID,
	}

	report, err := ths.repo.StoreReportComment(newReport)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_COMMENT_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf(constant.ERROR_REPORT_FAILED, "comment", req.CommentID),
			}
		}
	}

	return report.ToReportCommentResponse(), nil
}
