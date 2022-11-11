
BEGIN;

CREATE TABLE IF NOT EXISTS "auth".users (
    "id"            bigserial PRIMARY KEY,
    "email"         varchar(100) UNIQUE NOT NULL,
    "password"      varchar(100) NOT NULL,
    "first_name"    varchar(50),
    "last_name"     varchar(50),
    "is_active"     boolean NOT NULL,
    "secret_key"    varchar(50) NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "auth".profiles (
    "id"            bigserial PRIMARY KEY,
	"job"           varchar(60),
	"activities"    text,
    "photo"         varchar,
	"user_id"       bigint NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_profile_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE    
);

COMMENT ON TABLE "auth".users       IS 'User account basic information';
COMMENT ON TABLE "auth".profiles    IS 'User profile detail information';

COMMIT;