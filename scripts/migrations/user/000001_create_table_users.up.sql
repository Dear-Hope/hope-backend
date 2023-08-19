CREATE TABLE IF NOT EXISTS "user".users (
    "id"                    bigserial PRIMARY KEY,
    "email"                 varchar(100) UNIQUE NOT NULL,
    "password"              varchar(100) NOT NULL,
    "name"                  varchar(50),
    "alias"                 varchar(50),
    "is_verified"           boolean NOT NULL,
    "secret_key"            text,
    "photo"                 text,
    "total_audio_played"    int NOT NULL,
    "total_time_played"     int NOT NULL,
    "longest_streak"        int NOT NULL,
    "created_at"            timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"            timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"            boolean NOT NULL DEFAULT false
);
