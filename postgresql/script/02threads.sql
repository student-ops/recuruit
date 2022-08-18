\c recuruit

CREATE TABLE IF NOT EXISTS recuruitschema.threads(
    title varchar(50),
    userid varchar(50),
    datecreated timestamp,
    lang int,
    detail varchar(300),
    PRIMARY KEY (title,datecreated)
);
GRANT ALL PRIVILEGES ON recuruitschema.threads TO recuruit;
-- insertのところにもschema名を付ける必要がある
INSERT INTO recuruitschema.threads values(
    'first thread',
    'fefe',
    now(),
    1,
    'golang develop ment!'
);
INSERT INTO recuruitschema.threads values(
    'js dev',
    'rakky',
    now(),
    1,
    'js develop ment!'
);
