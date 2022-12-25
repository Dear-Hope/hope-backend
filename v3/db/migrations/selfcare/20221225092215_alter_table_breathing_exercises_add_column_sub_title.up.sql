BEGIN;

ALTER TABLE "selfcare".breathing_exercises 
    ADD COLUMN sub_title varchar NOT NULL DEFAULT '';

UPDATE "selfcare".breathing_exercises AS be 
    SET sub_title = u.sub_title
    FROM (
        VALUES
        (1, 'Tenangkan diri dan meringankan stress'),
        (2, 'Seimbangkan pikiran dan tubuh'),
        (3, 'Kembalikan fokus dan kurangi gugup'),
        (4, 'Tingkatkan energi dan kewaspadaan'),
        (5, 'Tenangkan pikiran dan lepaskan beban')
    ) AS u(id, sub_title)
    WHERE be.id = u.id;

COMMIT;