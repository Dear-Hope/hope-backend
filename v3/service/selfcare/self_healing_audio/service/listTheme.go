package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"net/http"
)

func (ths *service) ListTheme() ([]model.SelfHealingAudioThemeListResponse, *model.ServiceError) {
	selfHealingAudioThemes, err := ths.repo.GetAllTheme()
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_SELF_HEALING_AUDIO_THEMES_FAILED),
		}
	}

	return selfHealingAudioThemes.ToListSelfHealingAudioThemeResponse(), nil
}
