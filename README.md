## 💬 Usage

Costa は開発者向けの SNS です。
アカウントを作成し、プロジェクトメンバーの募集、メッセージのやり取りが可能です。

## 👉 Let's try 　

ログイン>ID:test PW:test でゲストユーザーとしてログインできます。

## 使用技術

Go docker aws postgres nginx

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
| ------------ | ------------------------- |
| user id      |                           |
| login id     | autoinclement, Primarykey |
| user name    |                           |
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
