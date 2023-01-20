BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".expert_topics (
    "id"            bigserial PRIMARY KEY,
    "expert_id"     bigint NOT NULL,
    "topic_id"      bigint NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    "is_deleted"    boolean NOT NULL DEFAULT false
);

COMMIT;