BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".posts (
    "id"                bigserial PRIMARY KEY,
    "author_id"         bigint NOT NULL,
    "content"           text NOT NULL,
    "is_deleted"        boolean NOT NULL DEFAULT false,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_post_author
        FOREIGN KEY (author_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_author_post" 
    ON "storyroom".posts USING btree ("author_id");

CREATE TABLE IF NOT EXISTS "storyroom".likes (
    "id"            bigserial PRIMARY KEY,
    "user_id"       bigint NOT NULL,
    "post_id"       bigint NOT NULL,
    "is_deleted"    boolean NOT NULL DEFAULT false,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_like_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_like_post
        FOREIGN KEY (post_id)
            REFERENCES "storyroom".posts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_user_post_like" 
    ON "storyroom".likes USING btree ("user_id", "post_id");    

COMMIT;