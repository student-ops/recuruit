CREATE DATABASE recuruit;
\c recuruit

-- ここがSHCEMAになってた

CREATE ROLE recuruit WITH LOGIN PASSWORD 'recuruit';
GRANT ALL PRIVILEGES ON DATABASE recuruit TO recuruit;