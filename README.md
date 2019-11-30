# README

- https://api-dot-soccer-250403.appspot.com/GetArticle/81

## 流れ
### 準備

1. `POST` id とpass
2. id とpassを検証
3. `Token or err`を返す

### 使用

1. `get soccerapi.exanple.com/GetArticle/81` + Token
2. 呼出元のドメインチェック（ホワイトリスト）
3. トークンの検証
4. Destの書き換え `soccerapi.exanple.com` => `api-dot-soccer-250403.appspot.com`

## todo

- auth / token class
  - cahce dataloader
  - 発行
  - 照合: bool

- クライアント（呼出元）
  - ドメインのホワイトリスト

- account / アカウント
  - id
  - password

- route
  - メモリキャッシュ
