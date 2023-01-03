BEGIN;

CREATE TABLE IF NOT EXISTS "auth".blocked_users (
    "id"                bigserial PRIMARY KEY,
	"user_id"           bigint NOT NULL,
	"blocked_user_id"   bigint NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now()),
    "updated_at"        timestamptz NOT NULL DEFAULT (now()),
    "is_deleted"        boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_blocking_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_blocked_user
        FOREIGN KEY (blocked_user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE            
);

CREATE UNIQUE INDEX IF NOT EXISTS "idx_blocked_user" 
    ON "auth".blocked_users USING btree ("user_id", "blocked_user_id"); 

COMMIT;