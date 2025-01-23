-- ユーザーに必要な権限を付与
CREATE USER IF NOT EXISTS 'twitter_user'@'%' IDENTIFIED BY 'twitter_password';
GRANT ALL PRIVILEGES ON twitter_clone.* TO 'twitter_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES; 