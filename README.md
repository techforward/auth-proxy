# README

- https://api-dot-soccer-250403.appspot.com/GetArticle/81
- ローカルでサブドメイン捌くのは`lvh.me`

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

## Check
```
curl -H "Origin: http://example.com" soccerapi.lvh.me/GetArticle/81
```

## todo

- dbなしで、決めうち
  - id: test
  - pass: test

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

## Ref

- https://github.com/rs/cors
- https://github.com/justinas/alice
- https://github.com/dgrijalva/jwt-go