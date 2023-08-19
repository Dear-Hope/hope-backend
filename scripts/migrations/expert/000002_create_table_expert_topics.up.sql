CREATE TABLE IF NOT EXISTS "expert".expert_topics (
    "id"            bigserial PRIMARY KEY,
    "expert_id"     bigint NOT NULL,
    "topic_id"      bigint NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"    boolean NOT NULL DEFAULT false
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_expert_topic" ON "expert".expert_topics USING btree("expert_id", "topic_id");
