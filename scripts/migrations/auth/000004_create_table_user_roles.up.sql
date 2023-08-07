CREATE TABLE IF NOT EXISTS "auth".user_roles(
    "id"            bigserial PRIMARY KEY,
    "user_id"       bigint NOT NULL,
    "role_name"          varchar(100) NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "is_deleted"    boolean NOT NULL DEFAULT false,
    CONSTRAINT fk_role_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    CONSTRAINT fk_role_name
        FOREIGN KEY (role_name)
            REFERENCES roles(name)
            ON UPDATE CASCADE
            ON DELETE CASCADE
);
