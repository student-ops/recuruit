CREATE DATABASE recuruit;
\c recuruit

CREATE SHCEMA recuruitschema;

CREATE ROLE recuruit WITH LOGIN PASSWORD 'recuruit';
GRANT ALL PRIVILEGES ON SCHEMA recuruitschema TO recuruit;