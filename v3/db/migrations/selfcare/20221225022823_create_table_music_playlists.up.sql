BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".music_playlists (
    "id"                bigserial PRIMARY KEY,
    "mood_id"           bigint NOT NULL,
    "title"             varchar(100) NOT NULL,
    "image_url"         text NOT NULL,
    "playlist_url"      text NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_mood
        FOREIGN KEY (mood_id)
            REFERENCES "moodtracker".moods(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

COMMIT;