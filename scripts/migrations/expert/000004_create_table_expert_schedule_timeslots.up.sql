CREATE TABLE IF NOT EXISTS "expert".expert_schedule_timeslots (
    "id"            bigserial PRIMARY KEY,
    "schedule_id"   bigint NOT NULL,
    "start_time"    varchar(25) NOT NULL,
    "end_time"      varchar(25) NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_schedule_timeslots
        FOREIGN KEY (schedule_id)
            REFERENCES "expert".expert_schedules(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);
