BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".needs (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(50) NOT NULL,    
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "selfcare".movies (
    "id"                bigserial PRIMARY KEY,
    "title"             varchar(255) NOT NULL,
    "year"              int not null,
    "country"           varchar(50) NOT NULL,
    "genres"            text[] NOT NULL,
    "description"       text NOT NULL,
    "poster_link"       text NOT NULL,
    "trailer_link"      text NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "selfcare".movie_mood_details (
    "id"                bigserial PRIMARY KEY,
    "movie_id"          bigint NOT NULL,
    "mood_id"           bigint NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_movie
        FOREIGN KEY (movie_id)
            REFERENCES "selfcare".movies(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_mood
        FOREIGN KEY (mood_id)
            REFERENCES "moodtracker".moods(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_movie_mood_id" 
    ON "selfcare".movie_mood_details USING btree ("movie_id");

CREATE TABLE IF NOT EXISTS "selfcare".movie_need_details (
    "id"                bigserial PRIMARY KEY,
    "movie_id"          bigint NOT NULL,
    "need_id"           bigint NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_movie
        FOREIGN KEY (movie_id)
            REFERENCES "selfcare".movies(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_need
        FOREIGN KEY (need_id)
            REFERENCES "selfcare".needs(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_movie_need_id" 
    ON "selfcare".movie_need_details USING btree ("movie_id");

COMMIT;