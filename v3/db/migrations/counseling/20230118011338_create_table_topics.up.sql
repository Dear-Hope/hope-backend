BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".topics (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(50) NOT NULL,
    "image_url"     varchar NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

COMMIT;