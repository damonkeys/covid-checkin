-- migrate:up
ALTER TABLE users ADD COLUMN preferred_language varchar(3) NULL DEFAULT 'en';

-- migrate:down
ALTER TABLE users DROP COLUMN preferred_language;
