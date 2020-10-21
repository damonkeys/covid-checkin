-- migrate:up
ALTER TABLE checkins CHANGE chckr_uuid user_uuid varchar(36);
ALTER TABLE checkins CHANGE chckr_name user_name varchar(500);
ALTER TABLE checkins CHANGE chckr_phone user_phone varchar(100);
ALTER TABLE checkins CHANGE chckr_email user_email varchar(255);
ALTER TABLE checkins CHANGE chckr_street user_street varchar(500);
ALTER TABLE checkins CHANGE chckr_city user_city varchar(100);
ALTER TABLE checkins CHANGE chckr_country user_country varchar(100);
ALTER TABLE checkins CHANGE chckr_registered user_registered boolean;

-- migrate:down
ALTER TABLE checkins CHANGE user_uuid chckr_uuid varchar(36);
ALTER TABLE checkins CHANGE user_name chckr_name varchar(500);
ALTER TABLE checkins CHANGE user_phone chckr_phone varchar(100);
ALTER TABLE checkins CHANGE user_email chckr_email varchar(255);
ALTER TABLE checkins CHANGE user_street chckr_street varchar(500);
ALTER TABLE checkins CHANGE user_city chckr_city varchar(100);
ALTER TABLE checkins CHANGE user_country chckr_country varchar(100);
ALTER TABLE checkins CHANGE user_registered chckr_registered boolean;
