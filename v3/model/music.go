package model

type (
	MusicPlaylist struct {
		ID          uint   `db:"id"`
		Title       string `db:"title"`
		ImageURL    string `db:"image_url"`
		PlaylistURL string `db:"playlist_url"`
		Mood        string `db:"mood"`
	}

	MusicPlaylists []MusicPlaylist
)

func (ths MusicPlaylist) TableWithSchemaName() string {
	return `"selfcare".music_playlists`
}

func (ths MusicPlaylist) ToMusicPlaylistResponse() *MusicPlaylistResponse {
	return &MusicPlaylistResponse{
		ID:          ths.ID,
		Title:       ths.Title,
		ImageURL:    ths.ImageURL,
		PlaylistURL: ths.PlaylistURL,
		Mood:        ths.Mood,
	}
}

func (ths MusicPlaylists) ToListMusicPlaylistResponse() []MusicPlaylistResponse {
	res := make([]MusicPlaylistResponse, len(ths))
	for i, movie := range ths {
		res[i] = *movie.ToMusicPlaylistResponse()
	}

	return res
}

type (
	MusicPlaylistResponse struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		ImageURL    string `json:"imageUrl"`
		PlaylistURL string `json:"playlistUrl"`
		Mood        string `json:"mood"`
	}
)
