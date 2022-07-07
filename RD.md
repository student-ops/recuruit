## 作成　ファイル
html 
 - create :アカウント作成
 - login :ログイン
 - top : トップページ(ログイン前)
 - top_after : トップページ(ログイン後)



[top]
Ι
Ι―[login(データ引き出し)]――[top_after]
Ι     ↑　　　　　			Ι
Ι―[create(データ挿入)]　　　　　 	Ι―[my_page(データ引き出し)]
　　　　　　　　　　　　　　　　	Ι
						Ι―[recuruit(データ引き出し)]
						Ι
						Ι―[create_recuruit(データ挿入)]

### DB

user table
|colum | datail |
| ----- | -----|
|user id | Primary key|
|user name| |
|pass word | |
|date created | |
|date updated| |
|date deleted | |

authorization

|colum | datail |
| ----- | -----|
|user id | |
|login id | autoinclement, Primarykey|
|user name| ||
|date created | |
|date updated| |

