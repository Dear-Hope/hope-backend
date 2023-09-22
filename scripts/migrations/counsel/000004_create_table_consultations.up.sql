CREATE TABLE IF NOT EXISTS "counsel".consultations (
    "id"            bigserial PRIMARY KEY,
    "user_id"       bigint NOT NULL,
    "expert_id"     bigint NOT NULL,
    "type_id"       bigint NOT NULL,
    "booking_date"  varchar(15) NOT NULL,
    "start_time"    varchar(10) NOT NULL,
    "end_time"      varchar(10) NOT NULL,
    "status"        varchar(25) NOT NULL,
    "user_notes"    TEXT NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_consul_user
        FOREIGN KEY (user_id)
            REFERENCES "user".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_consul_expert_id
        FOREIGN KEY (expert_id)
            REFERENCES "expert".experts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_consul_type
        FOREIGN KEY (type_id)
            REFERENCES "counsel".schedule_types(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS "booking_idx"
    ON "consultations" USING btree(expert_id, booking_date, start_time, end_time);
