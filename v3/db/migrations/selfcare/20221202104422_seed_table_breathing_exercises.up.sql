BEGIN;

INSERT INTO "selfcare".breathing_exercises(id, mood_id, title, name, repetition, description, benefit, implementation) VALUES
    (
        1,
        1, 
        'Relaksasi', 
        'Extended Exhale (4-6)',
        10, 
        'Tarik napas: 4 detik\nHembuskan: 6 detik\n\nAmbil napas yang dalam melalui hidung dan salurkan udara ke perut.Hembuskan perlahan melalui mulut dengan mengerucutkan bibir, sebagaimana meniup lilin.', 
        'Membantu memperlambat detak jantung dan mengendurkan saraf yang berguna untuk relaksasi dan mengurangi stress.',
        'Lakukan ketika merasa cemas, kurang istirahat, atau banyak pikiran.'    
    ),
    (
        2,
        4, 
        'Percaya Diri', 
        'Equal Breathing (4-2-4)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 2 detik\nHembuskan: 4 detik\n\nAmbil napas yang dalam melalui hidung dan salurkan udara ke perut.\nTahan saat di puncak pengambilan napas. Hembuskan perlahan selaras dengan ambil napas.', 
        'Durasi yang sama dalam mengambil dan menghembuskan napas membuat pergerakan diafragma secara berirama memiliki efek keseimbangan pada tubuh yang membantu tubuh rileks dan fokus.',
        'Lakukan ketika mengawali hari untuk merelaksasikan pikiran dan tubuh, dan menjadi percaya diri.'
    ),
    (
        3,
        5, 
        'Konsentrasi', 
        'Box Breathing (4-4-4-4)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 4 detik\nHembuskan: 4 detik\nTahan nafas: 4 detik\n\nTarik napas perlahan melalui hidung. Perhatikan dada Anda mengembang dan suara napas Anda Jeda di bagian atas tarikan napas Anda. Alih-alih menahan napas, anggap itu sebagai mengistirahatkan napas Anda, untuk menghindari ketegangan sebelum Anda menghembuskan napas. Buang napas perlahan dan lepaskan semua udara di paru-paru Anda. Tahan lagi dan istirahat sebelum menarik napas berikutnya', 
        'Membantu mengatur suhu tubuh dan menurunkan tekanan darah, serta membantu menjadi lebih penuh perhatian dan santai.',
        'Lakukan ketika akan menjalani momen penting seperti ujian, presentasi, dan berbicara di depan umum.'
    ),
    (
        4,
        2, 
        'Rehat', 
        'Relaxing Breath (4-7-8)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 7 detik\nHembuskan: 8 detik\n\nTarik napas dalam-dalam melalui hidung untuk mengisi perut dengan udara. Menghembuskan napas perlahan melalui mulut Anda dengan bibir mengerucut, seolah-olah Anda sedang meniup lilin. Anda mungkin mendengar suara "Whussss" ketika Anda menghembuskan napas', 
        'Membantu dalam mengendurkan otot-otot Anda, memperlambat irama detak jantung Anda, dan menenangkan pikiran Anda.',
        'Lakukan ketika merasa banyak pikiran untuk membantu tidur dan istirahat lebih baik.'
    ),
    (
        5,
        3, 
        'Berenergi', 
        'Energizing Breath (4-2)',
        10, 
        'Tarik napas: 4 detik\nHembuskan: 2 detik\n\nTarik napas dalam-dalam melalui hidung untuk mengisi perut dengan udara\\nMenghembuskan napas dengan cepat melalui hidung', 
        'Membantu melepaskan adrenalin dan meningkatkan aliran oksigen untuk membantu Anda merasa lebih terjaga, berenergi, dan waspada.',
        'Lakukan ketika kurang bersemangat pada pagi hari dan mengatasi lelah di sore hari.'
    );

COMMIT;