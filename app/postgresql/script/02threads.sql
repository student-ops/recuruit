\c recuruit

CREATE TABLE IF NOT EXISTS threads(
    thread_id BIGSERIAL PRIMARY KEY,
    title varchar(50),
    userid BIGINT,
    datecreated timestamp,
    lang int,
    detail varchar(600)
);
GRANT ALL PRIVILEGES ON threads TO recuruit;
INSERT INTO threads values(
    default,
    'first thread',
    1,
    now(),
    1,
    'golang develop ment!'
);
INSERT INTO threads values(
    default,
    'js dev',
    2,
    now(),
    1,
    'js develop ment!'
);
