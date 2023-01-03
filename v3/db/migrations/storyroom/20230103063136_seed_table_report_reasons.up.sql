BEGIN;

INSERT INTO "storyroom".report_reasons(id, reason) VALUES
    (1, 'Postingan ini tidak sesuai dengan kategori'),
    (2, 'Anda tidak tertarik dengan postingan ini'),
    (3, 'Postingan ini mengandung sara'),
    (4, 'Diserang karena identitas saya'),
    (5, 'Diganggu atau diintimidasi dengan kekerasan');

COMMIT;