CREATE TABLE IF NOT EXISTS "expert".expert_timeslot_types (
    "id"            BIGSERIAL PRIMARY KEY,
    "timeslot_id"   BIGINT NOT NULL,
    "type_id"       BIGINT NOT NULL,
    CONSTRAINT fk_timeslot
      FOREIGN KEY (type_id)
          REFERENCES "expert".expert_schedule_timeslots(id)
          ON UPDATE CASCADE
          ON DELETE CASCADE,
    CONSTRAINT fk_schedule_type
      FOREIGN KEY (type_id)
          REFERENCES "counsel".schedule_types(id)
          ON UPDATE CASCADE
          ON DELETE CASCADE
);
