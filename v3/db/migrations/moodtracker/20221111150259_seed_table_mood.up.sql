BEGIN;

INSERT INTO "moodtracker".moods(id, name) VALUES
    (1, 'Sad'),
    (2, 'Restless'),
    (3, 'Flat'),
    (4, 'Happy'),
    (5, 'Gorgeous');

COMMIT;