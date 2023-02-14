BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".expert_schedules (
    "id"            bigserial PRIMARY KEY,
    "expert_id"     bigint NOT NULL,
    "start_at"      timestamp(6) NOT NULL,
    "end_at"        timestamp(6) NOT NULL,
    "is_booked"     boolean NOT NULL DEFAULT false,
    "created_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_expert_schedule
     FOREIGN KEY (expert_id)
         REFERENCES "counseling".experts(id)
         ON UPDATE CASCADE
         ON DELETE CASCADE
);

COMMIT;