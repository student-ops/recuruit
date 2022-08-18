CREATE DATABASE recuruit;
\c recuruit

-- ここがSHCEMAになってた
CREATE SCHEMA recuruitschema;

CREATE ROLE recuruit WITH LOGIN PASSWORD 'recuruit';
GRANT ALL PRIVILEGES ON SCHEMA recuruitschema TO recuruit;