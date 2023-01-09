BEGIN;

CREATE TABLE IF NOT EXISTS "storyroom".report_comments (
    "id"            bigserial PRIMARY KEY,
    "user_id"       bigint NOT NULL,
    "comment_id"    bigint NOT NULL,
    "reason_id"     bigint NOT NULL,
    "is_deleted"    boolean NOT NULL DEFAULT false,
    "created_at"    timestamptz NOT NULL DEFAULT (now()),
    "updated_at"    timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_reporting_user
        FOREIGN KEY (user_id)
            REFERENCES "auth".users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_reported_comment
        FOREIGN KEY (comment_id)
            REFERENCES "storyroom".comments(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_report_reason
        FOREIGN KEY (reason_id)
            REFERENCES "storyroom".report_reasons(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE        
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_reported_comment" 
    ON "storyroom".report_comments USING btree ("user_id", "comment_id"); 

COMMIT;