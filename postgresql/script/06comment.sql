\c recuruit

CREATE TABLE IF NOT EXISTS comment(
    threadid BIGINT,
    added TIMESTAMP,
    userid BIGINT,
    comment_content varchar(300),
    PRIMARY KEY (threadid,added)
);
GRANT ALL PRIVILEGES ON comment TO recuruit;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO recuruit;
INSERT INTO comment values(1,NOW(),3,'開発参加したいです。');
INSERT INTO comment values(1,NOW(),1,'github,ポートフォリオ等ありますか?');