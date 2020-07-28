-- migrate:up
ALTER TABLE users
ADD COLUMN active boolean NOT NULL DEFAULT false,
ADD COLUMN activationToken varchar(36) NULL DEFAULT NULL,
ADD COLUMN activationTokenCreation timestamp NULL DEFAULT NULL,
ADD COLUMN activeSince timestamp NULL DEFAULT NULL;

UPDATE users set active=true, activeSince=NOW();
-- migrate:down
ALTER TABLE users
DROP COLUMN active,
DROP COLUMN activationToken,
DROP COLUMN activationTokenCreation,
DROP COLUMN activeSince;
