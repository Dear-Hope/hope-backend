BEGIN;

ALTER TABLE "auth".users DROP COLUMN "secret_key";

COMMIT;