CREATE TABLE IF NOT EXISTS "counsel".time_periods (
    "id"            bigserial PRIMARY KEY,
    "start_time"    varchar(10) NOT NULL,
    "end_time"      varchar(10) NOT NULL,
    "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
    "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW()))
);

INSERT INTO "counsel".time_periods (id, start_time, end_time)
    VALUES (1, '00:00', '01:00'),
           (2, '01:00', '02:00'),
           (3, '02:00', '03:00'),
           (4, '03:00', '04:00'),
           (5, '04:00', '05:00'),
           (6, '05:00', '06:00'),
           (7, '06:00', '07:00'),
           (8, '07:00', '08:00'),
           (9, '08:00', '09:00'),
           (10, '09:00', '10:00'),
           (11, '10:00', '11:00'),
           (12, '11:00', '12:00'),
           (13, '12:00', '13:00'),
           (14, '13:00', '14:00'),
           (15, '14:00', '15:00'),
           (16, '15:00', '16:00'),
           (17, '16:00', '17:00'),
           (18, '17:00', '18:00'),
           (19, '18:00', '19:00'),
           (20, '19:00', '20:00'),
           (21, '20:00', '21:00'),
           (22, '21:00', '22:00'),
           (23, '22:00', '23:00'),
           (24, '23:00', '24:00');
