-- migrate:up
ALTER TABLE users
ADD COLUMN activation_token varchar(36) NULL DEFAULT NULL,
ADD COLUMN activation_token_creation timestamp NULL DEFAULT NULL,
ADD COLUMN active_since timestamp NULL DEFAULT NULL;

UPDATE users set activation_token = activationToken;
UPDATE users set activation_token_creation = activationTokenCreation;
UPDATE users set active_since=activeSince;

ALTER TABLE users
DROP COLUMN activationToken,
DROP COLUMN activationTokenCreation,
DROP COLUMN activeSince;

-- migrate:down
ALTER TABLE users
ADD COLUMN activationToken varchar(36) NULL DEFAULT NULL,
ADD COLUMN activationTokenCreation timestamp NULL DEFAULT NULL,
ADD COLUMN activeSince timestamp NULL DEFAULT NULL;

UPDATE users set activationToken = activation_token;
UPDATE users set activationTokenCreation = activation_token_creation;
UPDATE users set activeSince = active_since;

ALTER TABLE users
DROP COLUMN activation_token,
DROP COLUMN activation_token_creation,
DROP COLUMN active_since;
