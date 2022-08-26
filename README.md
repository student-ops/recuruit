## 作成　ファイル

html

- create :アカウント作成
- login :ログイン
- top : トップページ(ログイン前)
- top_after : トップページ(ログイン後)

### DB

user table
|colum | datail |
| ----- | -----|
|user id | varchar*50)|
|pass word | varchar*50)|
|date created |date time |
|date updated| date time|
|date deleted | date time|
mysql
create table

authorization

| colum        | datail                    |
| ------------ | ------------------------- | --- |
| user id      |                           |
| login id     | autoinclement, Primarykey |
| user name    |                           |     |
| date created |                           |
| date updated |                           |

post
|colum|type|
|- |-|
|title|varchar(50) primary|
|datacreated| varchar(50) primary|
|lang|int|
|detail|varchar(300)|

memo
