\c recuruit
-- UserValuesからuservaluesに変更
CREATE TABLE if not exists uservalues(
    userid BIGSERIAL,
    username VARCHAR(50),
    created timestamp
);
-- UserValuesからuservaluesに変更
GRANT ALL PRIVILEGES ON uservalues TO recuruit;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO recuruit;
insert into uservalues values(DEFAULT,'test',now());
insert into uservalues values(DEFAULT,'harry kane',now());
insert into uservalues values(DEFAULT,'moh salah',now());
insert into uservalues values(DEFAULT,'reo messi',now());