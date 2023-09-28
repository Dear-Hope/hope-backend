ALTER TABLE "counsel".consultations
    ALTER COLUMN booking_date TYPE DATE USING to_date(booking_date, 'YYYY-MM-DD');
