BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".categories (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(100) NOT NULL,    
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

COMMIT;