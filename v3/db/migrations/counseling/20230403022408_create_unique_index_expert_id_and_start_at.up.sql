BEGIN;

CREATE UNIQUE INDEX IF NOT EXISTS "idx_expert_id_start_at"
    ON "counseling".expert_schedules USING btree ("expert_id", "start_at");

COMMIT;