BEGIN;

CREATE UNIQUE INDEX IF NOT EXISTS "idx_expert_topic" ON "counseling".expert_topics USING btree("expert_id", "topic_id");

COMMIT;