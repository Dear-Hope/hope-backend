package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"database/sql"
	"errors"
	"net/http"
)

func (ths *service) GetAudio(id uint) (*model.SelfHealingAudioResponse, *model.ServiceError) {
	selfHealingAudio, err := ths.repo.GetAudioByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &model.ServiceError{
				Code: http.StatusNotFound,
				Err:  errors.New(constant.ERROR_SELF_HEALING_AUDIO_NOT_FOUND),
			}
		}
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_SELF_HEALING_AUDIO_FAILED),
		}
	}

	return selfHealingAudio.ToSelfHealingAudioResponse(), nil
}
