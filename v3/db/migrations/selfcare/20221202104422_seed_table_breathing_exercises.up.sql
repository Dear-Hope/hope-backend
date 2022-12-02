BEGIN;

INSERT INTO "selfcare".breathing_exercises(id, title, name, repetition, description, benefit) VALUES
    (
        1, 
        'Tenang', 
        'Extended Exhale (4-6)',
        10, 
        'Tarik napas: 4 detik\nHembuskan: 6 detik\n\nTepat digunakan ketika kamu merasa cemas, kurang istirahat, atau banyak pikiran.', 
        'Mengambil napas panjang secara perlahan dapat mengirimkan sinyal relaksasi ke tubuh dan mengurangi stres. Menghembuskan napas lebih lama dibandingkan dengan menghirup membantu memperlambat detak jantung dan mengendurkan saraf.'
    ),
    (
        2, 
        'Keseimbangan', 
        'Equal Breathing (4-2-4)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 2 detik\nHembuskan: 4 detik\n\nTepat digunakan untuk mengawali hari untuk merelaksasikan pikiran dan tubuh, dan menjadi percaya diri.', 
        'Durasi yang sama dalam mengambil napas dan menghembuskannya yang membuat pergerakan diafragma secara berirama. Saat menghirup, darah dipindahkan dari paru-paru; saat menghembuskan napas, darah bergerak menuju paru-paru.Hal ini memiliki efek keseimbangan pada tubuh yang membantu tubuh rileks dan fokus.'
    ),
    (
        3, 
        'Fokus', 
        'Box Breathing (4-4-4-4)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 4 detik\nHembuskan: 4 detik\nTahan nafas: 4 detik\n\nTepat digunakan ketika sedang menjalankan momen penting seperti ujian, presentasi, rapat, dan berbicara didepan umum.', 
        'Pernapasan dalam yang berirama membantu mengatur suhu tubuh dan menurunkan tekanan darah. Menahan napas memungkinkan karbon dioksida menumpuk di dalam darah; melepaskan napas dapat merangsang tubuh Anda untuk membantu menjadi lebih penuh perhatian dan santai.'
    ),
    (
        4, 
        'Beristirahat', 
        'Relaxing Breath (4-7-8)',
        10, 
        'Tarik napas: 4 detik\nTahan nafas: 7 detik\nHembuskan: 8 detik\n\nTepat digunakan ketika merasa kewalahan, atau membantu kamu untuk tidur lebih baik.', 
        'Latihan ini mendorong pikiran dan tubuh Anda untuk fokus pada pengaturan napas, bukan pada kekhawatiran orang lain. Pernapasan lambat dengan menahan napas lebih lama membantu mengendurkan otot-otot Anda, memperlambat irama detak jantung Anda, dan menenangkan pikiran Anda.'
    ),
    (
        5, 
        'Berenergi', 
        'Energizing Breath (4-2)',
        10, 
        'Tarik napas: 4 detik\nHembuskan: 2 detik\n\nBaik dilakukan ketika merasa lelah atau lambat pada pagi hari atau lelah di sore hari.', 
        'Memperpendek napas dan bernapas dengan cepat, dengan menghirup lebih lama dan menghembuskan napas lebih pendek, membantu melepaskan adrenalin dan meningkatkan aliran oksigen untuk membantu Anda merasa lebih terjaga, berenergi, dan waspada.'
    );

COMMIT;