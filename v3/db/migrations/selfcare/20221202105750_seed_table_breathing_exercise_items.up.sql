BEGIN;

INSERT INTO "selfcare".breathing_exercise_items(id, exercise_id, name, duration, type) VALUES
    (1, 1, 'Tarik Nafas', 4, 'BREATH_IN'),
    (2, 1, 'Hembuskan', 6, 'BREATH_OUT'),
    (3, 2, 'Tarik Nafas', 4, 'BREATH_IN'),
    (4, 2, 'Tahan', 2, 'BREATH_HOLD'),
    (5, 2, 'Hembuskan', 4, 'BREATH_OUT'),
    (6, 3, 'Tarik Nafas', 4, 'BREATH_IN'),
    (7, 3, 'Tahan', 4, 'BREATH_HOLD'),
    (8, 3, 'Hembuskan', 4, 'BREATH_OUT'),
    (9, 3, 'Tahan', 4, 'BREATH_HOLD'),
    (10, 4, 'Tarik Nafas', 4, 'BREATH_IN'),
    (11, 4, 'Tahan', 7, 'BREATH_HOLD'),
    (12, 4, 'Hembuskan', 8, 'BREATH_OUT'),
    (13, 5, 'Tarik Nafas', 4, 'BREATH_IN'),
    (14, 5, 'Hembuskan', 2, 'BREATH_OUT');

COMMIT;