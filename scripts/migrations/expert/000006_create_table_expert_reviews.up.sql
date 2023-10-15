CREATE TABLE IF NOT EXISTS "expert".expert_reviews (
    "id"            BIGSERIAL PRIMARY KEY,
    "expert_id"     BIGINT NOT NULL,
    "rating"        BIGINT NOT NULL,
    "review"        TEXT,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_expert
      FOREIGN KEY (expert_id)
          REFERENCES "expert".experts(id)
          ON UPDATE CASCADE
          ON DELETE CASCADE
);
