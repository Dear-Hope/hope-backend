BEGIN;

CREATE TABLE IF NOT EXISTS "counseling".schedule_types (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(10) NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (now()),
    "is_deleted"    boolean NOT NULL DEFAULT false
);

INSERT INTO "counseling".schedule_types (id, name)
VALUES (1, 'CHAT'),
       (2, 'VOICE'),
       (3, 'VIDEO');

COMMIT;