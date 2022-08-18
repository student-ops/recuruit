\c recuruit
-- UserValuesからuservaluesに変更
CREATE TABLE if not exists recuruitschema.uservalues(
    userid varchar(50),
    password varchar(50),
    created timestamp
);
-- UserValuesからuservaluesに変更
GRANT ALL PRIVILEGES ON recuruitschema.uservalues TO recuruit;
-- schema名が必要
insert into recuruitschema.uservalues values('fefefe','fefefe',now());
insert into recuruitschema.uservalues values('moh','salah',now());