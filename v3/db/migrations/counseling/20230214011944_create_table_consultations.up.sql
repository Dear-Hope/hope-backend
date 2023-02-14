BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".consultations (
    "id"            bigserial PRIMARY KEY,
    "schedule_id"   bigint NOT NULL,
    "type_id"       bigint NOT NULL,
    "user_id"       bigint NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_consul_schedule
        FOREIGN KEY (schedule_id)
        REFERENCES "counseling".expert_schedules(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_consul_type
        FOREIGN KEY (type_id)
        REFERENCES "counseling".schedule_types(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT fk_consul_user
        FOREIGN KEY (user_id)
        REFERENCES "auth".users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

COMMIT;