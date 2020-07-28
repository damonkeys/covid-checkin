create database if not exists test_monkey_auth;
create user if not exists 'auth_user'@'%' identified by '';
grant ALL PRIVILEGES on test_monkey_auth.* to 'auth_user'@'%';
FLUSH PRIVILEGES;
