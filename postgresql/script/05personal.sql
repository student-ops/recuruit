\c recuruit

CREATE TABLE IF NOT EXISTS personal(
    userid BIGSERIAL,
    password VARCHAR(50)
);

GRANT ALL PRIVILEGES ON personal TO recuruit;
insert into personal values(DEFAULT,'test');
insert into personal values(DEFAULT,'spaas');
insert into personal values(DEFAULT,'revapool');
insert into personal values(DEFAULT,'barselona');
