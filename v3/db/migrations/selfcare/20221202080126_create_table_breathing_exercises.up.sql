BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".breathing_exercises (
    "id"                bigserial PRIMARY KEY,
    "mood_id"           bigint NOT NULL,
    "title"             varchar(100) NOT NULL,
    "name"              varchar(100) NOT NULL,
    "repetition"        int NOT NULL,
    "description"       text NOT NULL,
    "benefit"           text NOT NULL,
    "implementation"    text NOT NULL,       
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_mood
        FOREIGN KEY (mood_id)
            REFERENCES "moodtracker".moods(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "selfcare".breathing_exercise_items (
    "id"            bigserial PRIMARY KEY,
    "exercise_id"   bigint NOT NULL,
    "name"          varchar(100) NOT NULL,
    "duration"      int NOT NULL,
    "type"           varchar(100) NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_exercise
        FOREIGN KEY (exercise_id)
            REFERENCES "selfcare".breathing_exercises(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

COMMIT;