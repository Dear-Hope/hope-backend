CREATE TABLE IF NOT EXISTS "auth".user_profiles (
    "id"                    bigserial PRIMARY KEY,
    "user_id"               bigint NOT NULL,
    "photo"                 text,
    "total_audio_played"    int NOT NULL,
    "total_time_played"     int NOT NULL,
    "longest_streak"        int NOT NULL,
    "created_at"            timestamp(6) NOT NULL DEFAULT (NOW()),
    "updated_at"            timestamp(6) NOT NULL DEFAULT (NOW()),
    "is_deleted"            boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_profile_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);
