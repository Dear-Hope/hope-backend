package service

import (
	"HOPE-backend/v3/constant"
	"HOPE-backend/v3/model"
	"errors"
	"fmt"
	"net/http"
)

func (ths *service) ListAudioByMood(moodID, userID uint) ([]model.SelfHealingAudioListResponse, *model.ServiceError) {
	audios, err := ths.repo.GetAllAudioByMoodID(moodID)
	if err != nil {
		return nil, &model.ServiceError{
			Code: http.StatusInternalServerError,
			Err:  errors.New(constant.ERROR_GET_LIST_SELF_HEALING_AUDIO_THEMES_FAILED),
		}
	}

	var response []model.SelfHealingAudioListResponse
	for _, audio := range audios {
		lastOrderPlayed, err := ths.repo.GetLastOrderByThemeID(audio.ThemeID, userID)
		if err != nil {
			return nil, &model.ServiceError{
				Code: http.StatusInternalServerError,
				Err:  fmt.Errorf(constant.ERROR_GET_LAST_PLAYED_FAILED, "self healing audio"),
			}
		}
		response = append(response, audio.ToListItemSelfHealingAudioResponse(lastOrderPlayed+1))
	}

	return response, nil
}
