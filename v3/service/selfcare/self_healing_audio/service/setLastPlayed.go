package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func (ths *service) SetLastPlayed(userID uint, req model.SelfHealingAudioHistoryRequest) *model.ServiceError {
	selfHealingAudio, err := ths.repo.GetAudioByID(req.AudioID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_SELF_HEALING_AUDIO_NOT_FOUND),
			}
		}
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_SELF_HEALING_AUDIO_FAILED),
		}
	}

	_, err = ths.repo.StoreHistory(model.SelfHealingAudioHistory{
		AudioID: req.AudioID,
		ThemeID: selfHealingAudio.ThemeID,
		UserID:  userID,
		Order:   selfHealingAudio.Order,
	})
	if err != nil {
		return &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  fmt.Errorf(constant.ERROR_SET_LAST_PLAYED_FAILED, "self healing audio"),
		}
	}

	return nil
}
