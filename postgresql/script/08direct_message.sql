\c recuruit

CREATE TABLE IF NOT EXISTS direct_message(
    follow BIGINT,
    followed BIGINT,
    message varchar(300)
);

GRANT ALL PRIVILEGES ON direct_message TO recuruit;


INSERT INTO direct_message values(2,1,'こんにちは開発参加しませんか?');
INSERT INTO direct_message values(3,1,'こんにちはバックエンド開発参加しませんか?');
INSERT INTO direct_message values(4,1,'こんにちはアプリケーション開発参加しませんか?');
INSERT INTO direct_message values(1,2,'是非よろしくおねがいします');