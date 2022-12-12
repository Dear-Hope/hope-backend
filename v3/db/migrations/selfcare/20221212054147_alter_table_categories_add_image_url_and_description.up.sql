BEGIN;

ALTER TABLE "selfcare".categories 
    ADD COLUMN image_url varchar NOT NULL DEFAULT '',
    ADD COLUMN description varchar NOT NULL DEFAULT '';

UPDATE "selfcare".categories AS cat 
    SET image_url = u.image_url, description = u.description
    FROM (
        VALUES
        (1, 'https://res.cloudinary.com/shirotama/image/upload/v1670824810/image/selfcare/category/movie_mnebro.png', 'Kenali Film dan Series yang sesuai dengan kondisimu'),
        (2, 'https://res.cloudinary.com/shirotama/image/upload/v1670824811/image/selfcare/category/breath_exercise_jsdlzu.png', 'Eksplorasi Teknik Pernapasan yang sesuai dengan kebutuhanmu'),
        (3, 'https://res.cloudinary.com/shirotama/image/upload/v1670824811/image/selfcare/category/audio_self_healing_fapqxq.png', 'Dengarkan berbagai Audio Self-Healing yang cocok dengan kondisimu'),
        (4, 'https://res.cloudinary.com/shirotama/image/upload/v1670824810/image/selfcare/category/music_wzq8zf.png', 'Pilih Album Musik yang sesuai dengan kondisimu')
    ) AS u(id, image_url, description)
    WHERE cat.id = u.id;

COMMIT;