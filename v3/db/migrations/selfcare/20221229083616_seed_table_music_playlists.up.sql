BEGIN;

INSERT INTO "selfcare".music_playlists(id, mood_id, title, image_url, playlist_url) VALUES
    (1, 5, 'Tetap bahagia', 'https://images.unsplash.com/photo-1530021232320-687d8e3dba54?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=774&q=80', 'https://open.spotify.com/playlist/5ZwHab3h3vbRN012IbUg7t?si=791bdbe02b534bc0'),
    (2, 4, 'Mari berdansa', 'https://images.unsplash.com/photo-1665686377065-08ba896d16fd?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80', 'https://open.spotify.com/playlist/4z2zJ3t0ytaZUM5z4FjnHn?si=5ead8508983b4695'),
    (3, 3, 'Warnai harimu', 'https://images.unsplash.com/photo-1438761681033-6461ffad8d80?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80', 'https://open.spotify.com/playlist/3dg12IaxJzO7iH57yN4JIo?si=20c8913aa6ca41da'),
    (4, 1, 'Lebih tenang', 'https://images.unsplash.com/photo-1607688387751-c1e95ae09a42?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=774&q=80', 'https://open.spotify.com/playlist/3zCKAFCI1PEqKfWFThh9Nk?si=f9a6ef96434c431b'),
    (5, 2, 'Ayo bangkit!', 'https://images.unsplash.com/photo-1554188572-9d184b57d8e2?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1674&q=80', 'https://open.spotify.com/playlist/0jpAH5YTfHeMs6fMVsJKfP?si=3b5d671f80f2419d');

COMMIT;