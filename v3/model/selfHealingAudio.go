package model

import "time"

type (
	SelfHealingAudioTheme struct {
		ID              uint   `db:"id"`
		Title           string `db:"title"`
		Description     string `db:"description"`
		ImageURL        string `db:"image_url"`
		TotalAudio      int    `db:"total_audio"`
		LastPlayedOrder int    `db:"last_played"`

		Playlist SelfHealingAudios
	}

	SelfHealingAudioThemes []SelfHealingAudioTheme

	SelfHealingAudio struct {
		ID           uint   `db:"id"`
		ThemeID      uint   `db:"theme_id"`
		Title        string `db:"title"`
		ImageURL     string `db:"image_url"`
		AudioURL     string `db:"audio_url"`
		Description  string `db:"description"`
		Benefit      string `db:"benefit"`
		ScriptWriter string `db:"script_writer"`
		VoiceOver    string `db:"voice_over"`
		Duration     int    `db:"duration"`
		Order        int    `db:"order"`

		Subtitles SelfHealingAudioSubtitles
	}

	SelfHealingAudios []SelfHealingAudio

	SelfHealingAudioSubtitle struct {
		ID      uint   `db:"id"`
		AudioID uint   `db:"audio_id"`
		Text    string `db:"text"`
		Start   string `db:"start"`
		Order   int    `db:"order"`
	}

	SelfHealingAudioSubtitles []SelfHealingAudioSubtitle

	SelfHealingAudioHistory struct {
		ID        uint      `db:"id"`
		AudioID   uint      `db:"audio_id"`
		ThemeID   uint      `db:"theme_id"`
		UserID    uint      `db:"user_id"`
		Order     int       `db:"audio_order"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func (ths SelfHealingAudioTheme) TableWithSchemaName() string {
	return `"selfcare".self_healing_audio_themes`
}

func (ths SelfHealingAudio) TableWithSchemaName() string {
	return `"selfcare".self_healing_audios`
}

func (ths SelfHealingAudioSubtitle) TableWithSchemaName() string {
	return `"selfcare".self_healing_audio_subtitles`
}

func (ths SelfHealingAudioHistory) TableWithSchemaName() string {
	return `"selfcare".self_healing_audio_histories`
}

type (
	SelfHealingAudioThemeListResponse struct {
		ID            uint   `json:"id"`
		Title         string `json:"title"`
		Description   string `json:"description"`
		ImageURL      string `json:"imageUrl"`
		PlaylistCount int    `json:"playlistCount"`
	}

	SelfHealingAudioThemeResponse struct {
		ID       uint                           `json:"id"`
		Title    string                         `json:"title"`
		Playlist []SelfHealingAudioListResponse `json:"playlist"`
	}

	SelfHealingAudioListResponse struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		Creator  string `json:"creator"`
		ImageURL string `json:"imageUrl"`
		Duration int    `json:"durationSeconds"`
		IsLocked bool   `json:"isLocked"`
	}

	SelfHealingAudioResponse struct {
		ID           uint   `json:"id"`
		Title        string `json:"title"`
		ScriptWriter string `json:"scriptWriterName"`
		VoiceOver    string `json:"voiceOverName"`
		AudioURL     string `json:"audioUrl"`
		ImageURL     string `json:"backgroundImageUrl"`
		Description  string `json:"audioDesc"`
		Benefit      string `json:"benefitDesc"`

		Subtitles []SelfHealingAudioSubtitleResponse
	}

	SelfHealingAudioSubtitleResponse struct {
		ID    uint   `json:"id"`
		Text  string `json:"text"`
		Start string `json:"start"`
	}

	SelfHealingAudioHistoryRequest struct {
		AudioID uint `json:"audioID"`
	}
)

func (ths SelfHealingAudioThemes) ToListSelfHealingAudioThemeResponse() []SelfHealingAudioThemeListResponse {
	items := make([]SelfHealingAudioThemeListResponse, len(ths))
	for i, theme := range ths {
		items[i] = SelfHealingAudioThemeListResponse{
			ID:            theme.ID,
			Title:         theme.Title,
			Description:   theme.Description,
			ImageURL:      theme.ImageURL,
			PlaylistCount: theme.TotalAudio,
		}
	}

	return items
}

func (ths SelfHealingAudioTheme) ToSelfHealingAudioThemeResponse() *SelfHealingAudioThemeResponse {
	return &SelfHealingAudioThemeResponse{
		ID:       ths.ID,
		Title:    ths.Title,
		Playlist: ths.Playlist.ToListSelfHealingAudioResponse(ths.LastPlayedOrder + 1),
	}
}

func (ths SelfHealingAudios) ToListSelfHealingAudioResponse(lastPlayedOrder int) []SelfHealingAudioListResponse {
	items := make([]SelfHealingAudioListResponse, len(ths))
	for i, audio := range ths {
		items[i] = SelfHealingAudioListResponse{
			ID:       audio.ID,
			Title:    audio.Title,
			Creator:  audio.ScriptWriter,
			ImageURL: audio.ImageURL,
			Duration: audio.Duration,
			IsLocked: audio.Order > lastPlayedOrder,
		}
	}

	return items
}

func (ths SelfHealingAudio) ToSelfHealingAudioResponse() *SelfHealingAudioResponse {
	return &SelfHealingAudioResponse{
		ID:           ths.ID,
		Title:        ths.Title,
		ImageURL:     ths.ImageURL,
		AudioURL:     ths.AudioURL,
		ScriptWriter: ths.ScriptWriter,
		VoiceOver:    ths.VoiceOver,
		Description:  ths.Description,
		Benefit:      ths.Benefit,
		Subtitles:    ths.Subtitles.ToListSelfHealingAudioSubtitleResponse(),
	}
}

func (ths SelfHealingAudioSubtitles) ToListSelfHealingAudioSubtitleResponse() []SelfHealingAudioSubtitleResponse {
	items := make([]SelfHealingAudioSubtitleResponse, len(ths))
	for i, sub := range ths {
		items[i] = SelfHealingAudioSubtitleResponse{
			ID:    sub.ID,
			Text:  sub.Text,
			Start: sub.Start,
		}
	}

	return items
}
