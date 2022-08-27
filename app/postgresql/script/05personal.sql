\c recuruit

CREATE TABLE IF NOT EXISTS personal(
    userid BIGINT PRIMARY KEY,
    password VARCHAR(50)
);

GRANT ALL PRIVILEGES ON personal TO recuruit;
insert into personal values(1,'spaas');
insert into personal values(2,'revapool');