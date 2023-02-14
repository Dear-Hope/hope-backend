package constant

var (
	MAP_CATEGORY_TABLENAME map[string]string = map[string]string{
		"Sinema":             `"selfcare".movies`,
		"Latihan Pernapasan": `"selfcare".breathing_exercises`,
		"Audio Self-Healing": `"selfcare".self_healing_audios`,
		"Musik":              `"selfcare".music_playlists`,
	}
	CATEGORY_BREATHING_EXERCISE_ID = 2
	CATEGORY_AUDIO_SELF_HEALING_ID = 3
	CATEGORY_MUSIC_ID              = 4

	FormatDate = "2006-01-02"
	HourInSec  = 3600
)
