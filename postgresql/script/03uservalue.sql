\c recuruit
CREATE TABLE  if not exists recuruitschema.UserValues(
    userid varchar(50),
    password varchar(50),
    created timestamp
);
GRANT ALL PRIVILEGES ON recuruitschema.UserValues TO recuruit;
insert into uservalues values('fefefe','fefefe',now());
insert into uservalues values('moh','salah',now());
