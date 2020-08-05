-- migrate:up
ALTER TABLE users
CHANGE COLUMN `activationToken` `activation_token` varchar(36),
CHANGE COLUMN `activationTokenCreation` `activation_token_creation` timestamp,
CHANGE COLUMN `activeSince` `active_since` timestamp;

-- migrate:down
ALTER TABLE users
CHANGE COLUMN `activation_token` `activationToken` varchar(36),
CHANGE COLUMN `activation_token_creation` `activationTokenCreation` timestamp,
CHANGE COLUMN `active_since` `activeSince` timestamp;
