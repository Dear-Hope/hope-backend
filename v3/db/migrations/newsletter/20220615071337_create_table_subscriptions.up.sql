BEGIN;

CREATE TABLE IF NOT EXISTS "newsletter".subscriptions (
    "id"            bigserial PRIMARY KEY,
    "email"         varchar(100) UNIQUE NOT NULL,
    "subscribed_at" bigint NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON TABLE "newsletter".subscriptions IS 'Newsletter subscription info';

COMMIT;