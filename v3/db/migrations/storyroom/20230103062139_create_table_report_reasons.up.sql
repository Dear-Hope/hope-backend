BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".report_reasons (
    "id"            bigserial PRIMARY KEY,
    "reason"        TEXT NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

COMMIT;