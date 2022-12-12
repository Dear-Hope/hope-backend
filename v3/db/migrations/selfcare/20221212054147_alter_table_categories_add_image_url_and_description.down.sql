BEGIN;

ALTER TABLE "selfcare".categories 
    DROP COLUMN IF EXISTS image_url,
    DROP COLUMN IF EXISTS description;

COMMIT;