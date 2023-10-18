## React + Go(gin)を使用したログインのサンプルプログラム

### 概要
　勉強代わりにフロントはReactで開発して、バックエンドはGoを使用して開発を実施してみる。なお、ログイン認証はJWTを使用してみる。


### 指定可能な環境変数について

|環境変数名|説明|初期値|備考|
|--|--|--|--|
|WEB_HOST|バックエンドのWEBサーバのホスト名|空白|
|WEB_PORT|バックエンドのWEBサーバのポート|8080|
|WEB_ALLOW_ORIGIN|WEBサーバのヘッダー情報のAccess-Control-Allow-Originの設定|*|
|DB_TYPE|(設定変更不可)SQLサーバの種類指定|sqlite|
|DB_FILE|SQLiteの保存ファイル名|test.db|
|JWT_TOKEN_HOUR_LIFESPAN|JWTトークンの有効期間|24|時間表示|
|JWT_API_SECRET|JWTトークンのシークレットキー|auth-go-sample|
|AUTH_PEPPER|パスワードを暗号化する際のペッパー文字列|auth-go-sample|