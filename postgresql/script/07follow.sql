\c recuruit

CREATE TABLE IF NOT EXISTS follow(
    follower BIGINT,
    followed BIGINT
);
GRANT ALL PRIVILEGES ON follow TO recuruit;

INSERT INTO follow values(1,2);
INSERT INTO follow values(2,3);
INSERT INTO follow values(2,4);
INSERT INTO follow values(1,3);
INSERT INTO follow values(2,1);
INSERT INTO follow values(3,2);

