BEGIN;

CREATE TABLE IF NOT EXISTS "selfcare".items (
    "id" bigserial PRIMARY KEY,
    "mood" varchar(50) NOT NULL,
    "type" varchar(50) NOT NULL,
    "title" varchar(100) NOT NULL,
    "link" text NOT NULL,
    "description" text,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON TABLE "selfcare".items IS 'Self care recommendation items';

COMMIT;