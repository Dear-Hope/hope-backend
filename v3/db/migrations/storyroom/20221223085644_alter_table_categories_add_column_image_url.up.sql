BEGIN;

ALTER TABLE "storyroom".categories ADD COLUMN image_url varchar NOT NULL DEFAULT '';

COMMIT;