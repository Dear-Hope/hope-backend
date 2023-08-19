CREATE TABLE IF NOT EXISTS "expert".expert_schedules (
    "id"            bigserial PRIMARY KEY,
    "expert_id"     bigint NOT NULL,
    "day"           varchar(25) NOT NULL,
    "is_active"     boolean NOT NULL default false,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_expert_schedule
        FOREIGN KEY (expert_id)
            REFERENCES "expert".experts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_expert_id_day" ON "expert".expert_schedules USING btree ("expert_id", "day");
