CREATE TABLE UserValues(
    userid varchar(50),
    password bytea,
    created timestamp
);

CREATE TABLE test2(
    userid varchar(50),
    password bytea,
    created timestamp
);
insert into UserValues values('hurry','757eb8c038',now());