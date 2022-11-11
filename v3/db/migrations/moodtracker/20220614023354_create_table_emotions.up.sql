BEGIN;

CREATE TABLE IF NOT EXISTS "moodtracker".moods (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(50) NOT NULL,    
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "moodtracker".emotions (
    "id"                bigserial PRIMARY KEY,
    "mood_id"           bigint NOT NULL,
    "time_frame"        varchar(50) NOT NULL,
    "scale"             int NOT NULL,
    "description"       text NOT NULL,
    "triggers"          text[] NOT NULL,
    "feelings"          text[] NOT NULL,
    "date"              bigint NOT NULL,
    "user_id"           bigint NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_emotion_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_emotion_mood
        FOREIGN KEY (mood_id)
            REFERENCES "moodtracker".moods(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE     
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_date_time_frame_user_id" 
    ON "moodtracker".emotions USING btree ("user_id", "time_frame", ((to_timestamp("date") AT TIME ZONE 'UTC')::DATE)); 

COMMIT;