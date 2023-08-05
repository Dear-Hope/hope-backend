CREATE TABLE IF NOT EXISTS "auth".users (
    "id"            bigserial PRIMARY KEY,
    "email"         varchar(100) UNIQUE NOT NULL,
    "password"      varchar(100) NOT NULL,
    "name"          varchar(50),
    "alias"         varchar(50),
    "is_verified"   boolean NOT NULL,
    "secret_key"    text,
    "created_at"    timestamp(6) NOT NULL DEFAULT (NOW()),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (NOW()),
    "is_deleted"    boolean NOT NULL DEFAULT false
);
