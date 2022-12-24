BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".self_healing_audio_themes (
    "id"            bigserial PRIMARY KEY,
    "title"         varchar(100) NOT NULL,
    "description"   text NOT NULL,
    "image_url"     text NOT NULL,       
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "selfcare".self_healing_audios (
    "id"            bigserial PRIMARY KEY,
    "theme_id"      bigint NOT NULL,
    "mood_id"       bigint NOT NULL,
    "title"         varchar(100) NOT NULL,
    "description"   text NOT NULL,
    "image_url"     text NOT NULL,
    "audio_url"     text NOT NULL,
    "benefit"       text NOT NULL,
    "script_writer" varchar(100) NOT NULL,      
    "voice_over"    varchar(100) NOT NULL,
    "duration"      int NOT NULL,
    "order"      int NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_theme
        FOREIGN KEY (theme_id)
            REFERENCES "selfcare".self_healing_audio_themes(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "selfcare".self_healing_audio_subtitles (
    "id"            bigserial PRIMARY KEY,
    "audio_id"      bigint NOT NULL,
    "text"          text NOT NULL,
    "start"         varchar(10) NOT NULL,
    "order"         int NOT NULL,       
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_audio
        FOREIGN KEY (audio_id)
            REFERENCES "selfcare".self_healing_audios(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

COMMIT;