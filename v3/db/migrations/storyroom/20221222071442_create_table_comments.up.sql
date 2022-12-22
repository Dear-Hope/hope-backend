BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".comments (
    "id"                bigserial PRIMARY KEY,
    "post_id"           bigint NOT NULL,
    "author_id"         bigint NOT NULL,
    "content"           text NOT NULL,
    "is_deleted"        boolean NOT NULL DEFAULT false,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_comment_author
        FOREIGN KEY (author_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_comment_post
        FOREIGN KEY (post_id)
            REFERENCES "storyroom".posts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE        
);

CREATE INDEX IF NOT EXISTS "idx_post_comment" 
    ON "storyroom".comments USING btree ("post_id");

COMMIT;