BEGIN;

DROP TABLE IF EXISTS "auth".blocked_users;

DROP INDEX IF EXISTS "idx_blocked_user";

COMMIT;