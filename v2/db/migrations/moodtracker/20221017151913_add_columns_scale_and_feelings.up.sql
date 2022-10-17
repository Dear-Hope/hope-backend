BEGIN;

ALTER TABLE "moodtracker".emotions
ADD feelings text[] NOT NULL,
ADD scale int NOT NULL;

COMMIT;