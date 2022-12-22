BEGIN;

DROP TABLE IF EXISTS "storyroom".posts;
DROP TABLE IF EXISTS "storyroom".likes;

DROP INDEX IF EXISTS "idx_author_post";
DROP INDEX IF EXISTS "idx_user_post_like";

COMMIT;