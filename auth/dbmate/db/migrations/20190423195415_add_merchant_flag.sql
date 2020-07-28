-- migrate:up
ALTER TABLE
  users
ADD
  merchant boolean DEFAULT false;
-- migrate:down
ALTER TABLE
  users DROP COLUMN IF EXISTS merchant;
