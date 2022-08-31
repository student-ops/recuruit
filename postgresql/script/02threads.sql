\c recuruit

CREATE TABLE IF NOT EXISTS threads(
    threadid BIGSERIAL ,
    title varchar(50),
    userid BIGINT,
    datecreated timestamp,
    lang int,
    detail varchar(600)
);

GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO recuruit;
GRANT ALL PRIVILEGES ON threads TO recuruit;
INSERT INTO threads values(
    default,
    'go言語でのバックエンド開発',
    1,
    now(),
    1,
    'go言語を使用したマイクロサービスアーキテクチャの開発です。'
);
INSERT INTO threads values(
    default,
    'javascriptフロントエンド開発です。',
    2,
    now(),
    1,
    'javascriptでのフロントエンド開発です。'
);
