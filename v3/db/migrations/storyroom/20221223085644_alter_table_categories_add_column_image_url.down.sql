BEGIN;

ALTER TABLE "selfcare".categories DROP COLUMN IF EXISTS image_url;

COMMIT;