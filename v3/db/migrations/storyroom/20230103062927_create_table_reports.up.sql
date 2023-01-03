BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".reports (
    "id"            bigserial PRIMARY KEY,
    "user_id"       bigint NOT NULL,
    "post_id"       bigint NOT NULL,
    "reason_id"     bigint NOT NULL,
    "is_deleted"    boolean NOT NULL DEFAULT false,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_reporting_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_reported_post
        FOREIGN KEY (post_id)
            REFERENCES "storyroom".posts(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_report_reason
        FOREIGN KEY (reason_id)
            REFERENCES "storyroom".report_reasons(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE        
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_reported_post" 
    ON "storyroom".likes USING btree ("user_id", "post_id"); 

COMMIT;