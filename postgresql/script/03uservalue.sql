\c recuruit
-- UserValuesからuservaluesに変更
CREATE TABLE if not exists uservalues(
    userid varchar(50),
    password varchar(50),
    created timestamp
);
-- UserValuesからuservaluesに変更
GRANT ALL PRIVILEGES ON uservalues TO recuruit;
-- schema名が必要
insert into uservalues values('fefefe','fefefe',now());
insert into uservalues values('moh','salah',now());