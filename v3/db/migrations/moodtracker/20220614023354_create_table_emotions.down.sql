BEGIN;

DROP INDEX IF EXISTS "idx_date_time_frame_user_id";
DROP TABLE IF EXISTS "moodtracker".emotions;
DROP TABLE IF EXISTS "moodtracker".moods;

COMMIT;