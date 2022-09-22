BEGIN;

ALTER TABLE "auth".users ADD COLUMN "secret_key" varchar(50) DEFAULT '' NOT NULL;

COMMIT;