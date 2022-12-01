BEGIN;

DROP TABLE IF EXISTS "selfcare".needs;
DROP TABLE IF EXISTS "selfcare".movies;
DROP TABLE IF EXISTS "selfcare".movie_mood_details;
DROP TABLE IF EXISTS "selfcare".movie_need_details;

DROP INDEX IF EXISTS "idx_movie_mood_id";
DROP INDEX IF EXISTS "idx_movie_need_id";

COMMIT;