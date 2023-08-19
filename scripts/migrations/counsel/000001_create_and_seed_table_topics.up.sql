CREATE TABLE IF NOT EXISTS "counsel".topics (
   "id"            bigserial PRIMARY KEY,
   "name"          varchar(50) NOT NULL,
   "image_url"     varchar NOT NULL,
   "created_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW())),
   "updated_at"    timestamp(6) NOT NULL DEFAULT (TIMEZONE('UTC', NOW()))
);

INSERT INTO "counsel".topics(id, name, image_url) VALUES
    (1, 'Teman', 'assets/images/mood_triggers/teman.png'),
    (2, 'Pendidikan', 'assets/images/mood_triggers/pendidikan.png'),
    (3, 'Keluarga', 'assets/images/mood_triggers/keluarga.png'),
    (4, 'Keuangan', 'assets/images/mood_triggers/keuangan.png'),
    (5, 'Spiritual', 'assets/images/mood_triggers/spiritual.png'),
    (6, 'Self Care', 'assets/images/mood_triggers/self_care.png'),
    (7, 'Asmara', 'assets/images/mood_triggers/pasangan.png'),
    (8, 'Pekerjaan', 'assets/images/mood_triggers/pekerjaan.png');
