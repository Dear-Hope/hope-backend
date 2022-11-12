BEGIN;

INSERT INTO "moodtracker".moods(id, name) VALUES
    (1, 'Angry'),
    (2, 'Sad'),
    (3, 'Flat'),
    (4, 'Happy'),
    (5, 'Gorgeous');

COMMIT;