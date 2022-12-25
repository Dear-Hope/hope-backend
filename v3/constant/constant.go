package constant

var (
	MAP_CATEGORY_TABLENAME map[string]string = map[string]string{
		"Sinema":             `"selfcare".movies`,
		"Latihan Pernafasan": `"selfcare".breathing_exercises`,
		"Audio Self-Healing": `"selfcare".self_healing_audios`,
		"Musik":              `"selfcare".music_playlists`,
	}
	CATEGORY_BREATHING_EXERCISE_ID = 2
	CATEGORY_AUDIO_SELF_HEALING_ID = 3
	CATEGORY_MUSIC_ID              = 4
)
