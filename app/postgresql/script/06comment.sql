\c recuruit

CREATE TABLE IF NOT EXISTS comment(
    thread_id BIGINT,
    comment_no SERIAL,
    userid BIGINT,
    comment_content varchar(300),
    PRIMARY KEY (thread_id,userid)
);
GRANT ALL PRIVILEGES ON comment TO recuruit;
INSERT INTO comment values(1,DEFAULT,3,'hey sapp');
INSERT INTO comment values(1,DEFAULT,1,'good');