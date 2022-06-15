BEGIN;

CREATE TABLE IF NOT EXISTS "moodtracker".emotions (
    "id" bigserial PRIMARY KEY,
    "mood" varchar(50) NOT NULL,
    "time_frame" varchar(50) NOT NULL,
    "description" text NOT NULL,
    "triggers" text[] NOT NULL,
    "date" bigint NOT NULL,
    "user_id" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_emotion_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE    
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_date_time_frame_user_id" 
    ON "moodtracker".emotions USING btree ("user_id", "time_frame", "date"); 

COMMIT;