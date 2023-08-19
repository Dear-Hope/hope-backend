CREATE TABLE IF NOT EXISTS "counsel".schedule_types (
    "id"            bigserial PRIMARY KEY,
    "name"          varchar(10) NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW()))
);

INSERT INTO "counsel".schedule_types (id, name)
VALUES (1, 'CHAT'),
       (2, 'VOICE'),
       (3, 'VIDEO');
