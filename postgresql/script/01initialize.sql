CREATE DATABASE recuruit;
\c recuruit


CREATE ROLE recuruit WITH LOGIN PASSWORD 'recuruit';
GRANT ALL PRIVILEGES ON DATABASE recuruit TO recuruit;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO recuruit;
