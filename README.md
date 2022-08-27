## Difference btween master

    potgres dockercompose の追加。統合docker composeを削除　app/queryのDB接続設定を変更。

## Postgres DB

user values
|colum | datail |
| ----- | -----|
|user id | bigint auto incremnt primarykey|
|user name | varchar(50)|
|date created |date time |
|-|-|

user profile

| colum           | type               |
| --------------- | ------------------ |
| userid          | bigint primary key |
| back end skill  | int (bit)          |
| front end skill | int (bit)          |
| infla skill     | int(bit)           |
| message         | varcahr(200)       |

personal
|colum | datail |
| ----- | -----|
|user id | bit int primarykey|
|pass word| string //plan bit//|
| mail address//paln | string|

thread
|colum|type|
|--|--|
|threadid| BIGSERIAL PRIMARY KEY|
|title|varchar(50) |
|datacreated| varchar(50)|
|user id|int|
|lang|int|
|detail |varchar(600)|

commnet
|colum|type|
|--|-|
|userid | -|
|thread name |varchar(50) primary key|
|thread created|varchar(50) primary key|
|id|varchar (50) primary key|
|date added|date time|
|content| varchar (400)|

memo
