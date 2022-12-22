BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".categories (
    "id"                bigserial PRIMARY KEY,
    "name"              varchar(50) NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "storyroom".category_posts (
    "id"                bigserial PRIMARY KEY,
    "post_id"           bigint NOT NULL,
    "category_id"       bigint NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id)
            REFERENCES "storyroom".categories(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_post
        FOREIGN KEY (post_id)
            REFERENCES "storyroom".posts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS "idx_post_categories" 
    ON "storyroom".category_posts USING btree ("post_id");

CREATE INDEX IF NOT EXISTS "idx_category_posts" 
    ON "storyroom".category_posts USING btree ("category_id");

COMMIT;