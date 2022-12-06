BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".self_healing_audio_histories (
    "id"            bigserial PRIMARY KEY,
    "theme_id"      bigint NOT NULL,
    "audio_id"      bigint NOT NULL,
    "user_id"       bigint NOT NULL,
    "audio_order"   int NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_theme
        FOREIGN KEY (theme_id)
            REFERENCES "selfcare".self_healing_audio_themes(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_audio
        FOREIGN KEY (audio_id)
            REFERENCES "selfcare".self_healing_audios(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_theme_audio_user" 
    ON "selfcare".self_healing_audio_histories USING btree ("user_id", "audio_id", "theme_id");

COMMIT;