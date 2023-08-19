CREATE TABLE IF NOT EXISTS "expert".experts (
    "id"                    bigserial PRIMARY KEY,
    "email"                 varchar(100) UNIQUE NOT NULL,
    "password"              varchar(100) NOT NULL,
    "name"                  varchar NOT NULL,
    "expertise"             varchar(50) NOT NULL,
    "rating"                decimal(2,1) NOT NULL DEFAULT 0,
    "price"                 int NOT NULL DEFAULT 0,
    "is_available"          boolean NOT NULL DEFAULT false,
    "title"                 varchar NOT NULL,
    "education"             text NOT NULL,
    "experience"            text NOT NULL,
    "photo"                 text,
    "bio"                   text,
    "created_at"            timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"            timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"            boolean NOT NULL DEFAULT false
);
