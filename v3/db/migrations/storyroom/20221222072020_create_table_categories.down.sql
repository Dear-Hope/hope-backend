BEGIN;

DROP TABLE IF EXISTS "storyroom".categories;
DROP TABLE IF EXISTS "storyroom".category_posts;

DROP INDEX IF EXISTS "idx_post_categories";
DROP INDEX IF EXISTS "idx_category_posts";

COMMIT;