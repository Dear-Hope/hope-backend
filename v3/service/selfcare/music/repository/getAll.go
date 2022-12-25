package repository

import (
	"HOPE-backend/v3/model"
	"HOPE-backend/v3/service/selfcare/filter"
	"log"
)

func (ths *repository) GetAll(f filter.ListMusic) (model.MusicPlaylists, error) {
	var (
		playlists model.MusicPlaylists
		where     string
		args      []interface{}
	)

	if f.MoodID > 0 {
		where = "AND mp.mood_id = $1"
		args = append(args, f.MoodID)
	}

	err := ths.db.Select(
		&playlists,
		`SELECT mp.id, mp.title, mp.image_url, mp.playlist_url, m.name as mood 
		FROM "selfcare".music_playlists mp, "moodtracker".moods m 
		WHERE mp.mood_id = m.id `+where,
		args...,
	)
	if err != nil {
		log.Printf("get all music playlists: %s", err.Error())
		return nil, err
	}

	return playlists, nil
}
