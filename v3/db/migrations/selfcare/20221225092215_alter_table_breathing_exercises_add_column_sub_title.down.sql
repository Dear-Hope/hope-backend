BEGIN;

ALTER TABLE "selfcare".breathing_exercises 
    DROP COLUMN IF EXISTS sub_title;

COMMIT;