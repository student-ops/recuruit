CREATE TABLE IF NOT EXISTS threads(
    title varchar(50),
    datecreated timestamp,
    lang int,
    detail varchar(300),
    PRIMARY KEY (title,datecreated)
);
INSERT INTO threads values(
    'first thread',
    now(),
    1,
    'golang develop ment!'
);