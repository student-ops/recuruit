## Difference btween master

    potgres dockercompose の追加。統合docker composeを削除　app/queryのDB接続設定を変更。

## Postgres DB

user values
|colum | datail |
| ----- | -----|
|user id | varchar*50)|
|pass word | varchar*50)|
|date created |date time |
|date updated| date time|
|date deleted | date time|
|-|-|

//authorization

| colum        | datail                    |
| ------------ | ------------------------- |
| user id      |                           |
| login id     | autoinclement, Primarykey |
| user name    |                           |
| date created |                           |
| date updated |                           |
| -            | -                         |

thread
|colum|type|
|- |-|
|title|varchar(50) primary|
|datacreated| varchar(50) primary|
|lang|int|
|detail|varchar(300)|
|-|-|

memo
