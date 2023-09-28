ALTER TABLE "counsel".consultations
    DROP COLUMN IF EXISTS counsel_notes,
    DROP COLUMN IF EXISTS  suggestion,
    DROP COLUMN IF EXISTS  activity_recommendation,
    DROP COLUMN IF EXISTS  document;
