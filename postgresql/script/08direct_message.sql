\c recuruit
DROP TABLE IF EXISTS direct_message;
CREATE TABLE IF NOT EXISTS direct_message(
    follow BIGINT,
    followed BIGINT,
    message varchar(300),
    time TIMESTAMP
);

GRANT ALL PRIVILEGES ON direct_message TO recuruit;


INSERT INTO direct_message values(2,1,'こんにちは開発参加しませんか?',now());
INSERT INTO direct_message values(2,1,'進捗いかがでしょうか',now());
INSERT INTO direct_message values(3,1,'こんにちはバックエンド開発参加しませんか?',now());
INSERT INTO direct_message values(1,2,'是非よろしくおねがいします',now());
INSERT INTO direct_message values(2,1,'詳細以下になります。',now());
INSERT INTO direct_message values(4,1,'こんにちはアプリケーション開発参加しませんか?',now());