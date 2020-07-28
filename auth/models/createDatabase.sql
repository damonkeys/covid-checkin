create database if not exists ch3ck1n;
create user if not exists 'ch3ck1n_user'@'%' identified by '==>';
grant ALL PRIVILEGES on ch3ck1n.* to 'ch3ck1n_user'@'%';
FLUSH PRIVILEGES;
