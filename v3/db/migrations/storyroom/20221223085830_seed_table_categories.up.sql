BEGIN;

INSERT INTO "storyroom".categories(id, name, image_url) VALUES
    (1, 'Teman', 'assets/images/mood_triggers/teman.png'),
    (2, 'Pendidikan', 'assets/images/mood_triggers/pendidikan.png'),
    (3, 'Keluarga', 'assets/images/mood_triggers/keluarga.png'),
    (4, 'Keuangan', 'assets/images/mood_triggers/keuangan.png'),
    (5, 'Spiritual', 'assets/images/mood_triggers/spiritual.png'),
    (6, 'Self Care', 'assets/images/mood_triggers/self_care.png'),
    (7, 'Asmara', 'assets/images/mood_triggers/pasangan.png'),
    (8, 'Pekerjaan', 'assets/images/mood_triggers/pekerjaan.png');

COMMIT;