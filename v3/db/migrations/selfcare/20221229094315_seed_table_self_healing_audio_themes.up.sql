BEGIN;

INSERT INTO "selfcare".self_healing_audio_themes(id, title, description, image_url) VALUES
    (1, 'Beraktivitas', 'Mulai hari dengan penuh semangat dan berenergi', 'https://images.unsplash.com/photo-1540539234-c14a20fb7c7b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80'),
    (2, 'Istirahat', 'Tenangkan pikiran dan beristirahat dengan nyaman', 'https://images.unsplash.com/photo-1568617935424-49ab968826d7?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1470&q=80');
    
COMMIT;