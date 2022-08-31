\c recuruit

CREATE TABLE IF NOT EXISTS userprofile(
    userid BIGINT,
    --backend int,
    --frontend int,
    --infra int,
    profiletext VARCHAR(400)
);

GRANT ALL PRIVILEGES ON userprofile TO recuruit;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO recuruit;
INSERT INTO userprofile values (2,'hello form spaas');
INSERT INTO userprofile values (3,'hello form revapool');
INSERT INTO userprofile values (4,'hello from barserona');
