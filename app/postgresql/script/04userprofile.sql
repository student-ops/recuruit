\c recuruit

CREATE TABLE IF NOT EXISTS userprofile(
    userid BIGSERIAL PRIMARY KEY,
    backend int,
    frontend int,
    infra int,
    profiletext VARCHAR(400)
);

GRANT ALL PRIVILEGES ON userprofile TO recuruit;
INSERT INTO userprofile values (DEFAULT,0,0,0,'hello form spaas');
INSERT INTO userprofile values (DEFAULT,0,0,0,'hello form revapool');
INSERT INTO userprofile values (DEFAULT,0,0,0,'hello from barserona');