BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".experts (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar NOT NULL,
    "expertise"     varchar(50) NOT NULL,
    "rating"        decimal(2,1) NOT NULL DEFAULT 0,
    "price"         int NOT NULL DEFAULT 0,
    "is_available"  boolean NOT NULL DEFAULT false,
    "title"         varchar NOT NULL,
    "education"     text NOT NULL,
    "experience"    text NOT NULL,
    "photo"         text,
    "bio"           text,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    "is_deleted"    boolean NOT NULL DEFAULT false
);

COMMIT;