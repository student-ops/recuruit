\c recuruit

CREATE TABLE IF NOT EXISTS comment(
    thread_id BIGINT,
    added TIMESTAMP,
    userid BIGINT,
    comment_content varchar(300),
    PRIMARY KEY (thread_id,added)
);
GRANT ALL PRIVILEGES ON comment TO recuruit;
INSERT INTO comment values(1,NOW(),3,'hey sapp');
INSERT INTO comment values(1,NOW(),1,'good');