-- migrate:up
ALTER TABLE businesses MODIFY COLUMN name varchar(50);
ALTER TABLE business_infos MODIFY COLUMN description MEDIUMTEXT NULL;

ALTER TABLE businesses ADD COLUMN street varchar(50) NULL;
ALTER TABLE businesses ADD COLUMN zip varchar(10) NULL;
ALTER TABLE businesses ADD COLUMN city varchar(100) NULL;
ALTER TABLE businesses ADD COLUMN country varchar(60) NULL;


-- migrate:down
ALTER TABLE businesses MODIFY COLUMN name varchar(15);
ALTER TABLE business_infos MODIFY COLUMN description varchar(255);

ALTER TABLE businesses DROP COLUMN street;
ALTER TABLE businesses DROP COLUMN zip;
ALTER TABLE businesses DROP COLUMN city;
ALTER TABLE businesses DROP COLUMN country;

ALTER TABLE business_infos DROP COLUMN small_description;
