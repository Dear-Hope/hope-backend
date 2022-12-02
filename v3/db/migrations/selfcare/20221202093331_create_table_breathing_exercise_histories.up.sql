BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".breathing_exercise_histories (
    "id"            bigserial PRIMARY KEY,
    "exercise_id"   bigint NOT NULL,
    "user_id"       bigint NOT NULL,       
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_exercise
        FOREIGN KEY (exercise_id)
            REFERENCES "selfcare".breathing_exercises(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_exercise_user" 
    ON "selfcare".breathing_exercise_histories USING btree ("user_id", "exercise_id");

COMMIT;