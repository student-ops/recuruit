\c recuruit

CREATE TABLE IF NOT EXISTS threads(
    title varchar(50),
    userid varchar(50),
    datecreated timestamp,
    lang int,
    detail varchar(300),
    PRIMARY KEY (title,datecreated)
);
GRANT ALL PRIVILEGES ON threads TO recuruit;
-- insertのところにもschema名を付ける必要がある
INSERT INTO threads values(
    'first thread',
    'fefe',
    now(),
    1,
    'golang develop ment!'
);
INSERT INTO threads values(
    'js dev',
    'rakky',
    now(),
    1,
    'js develop ment!'
);
