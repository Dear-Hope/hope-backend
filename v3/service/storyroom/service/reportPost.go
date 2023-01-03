package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (ths *service) ReportPost(req model.ReportRequest, userID uint) (*model.ReportResponse, *model.ServiceError) {
	if req.PostID <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_REPORT_POST_FAILED, req.PostID),
		}
	}

	if req.ReasonID <= 0 {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_REPORT_POST_MINIMAL_ONE_REASON),
		}
	}

	newReport := model.Report{
		UserID:   userID,
		PostID:   req.PostID,
		ReasonID: req.ReasonID,
	}

	report, err := ths.repo.StoreReport(newReport)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &model.ServiceError{
				Code: http.StatusBadRequest,
				Err:  errors.New(constant.ERROR_COMMENT_ALREADY_EXISTS),
			}
		} else {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf(constant.ERROR_REPORT_POST_FAILED, req.PostID),
			}
		}
	}

	return report.ToReportResponse(), nil
}
