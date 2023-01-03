package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) ListReason() ([]model.ReportReasonResponse, *model.ServiceError) {
	reasons, err := ths.repo.GetAllReason()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_REASON_FAILED),
		}
	}

	return reasons.ToListReportReasonResponse(), nil
}
