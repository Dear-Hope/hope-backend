CREATE TABLE IF NOT EXISTS "auth".roles(
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(100) NOT NULL UNIQUE,
    "created_at"    timestamp(6) NOT NULL DEFAULT (NOW()),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (NOW()),
    "is_deleted"    boolean NOT NULL DEFAULT false
);
