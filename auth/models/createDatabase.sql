create database if not exists monkey_auth;
create user if not exists 'auth_user'@'%' identified by '';
grant ALL PRIVILEGES on monkey_auth.* to 'auth_user'@'%';
FLUSH PRIVILEGES;
