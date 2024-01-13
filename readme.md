## React + Go(gin)を使用したログインのサンプルプログラム

### 概要

勉強代わりにフロントはReactで開発して、バックエンドはGoを使用して開発を実施してみる。なお、ログイン認証はJWTを使用してみる。

### 指定可能な環境変数について

| 環境変数名              | 説明                                                       | 初期値         | 備考     |
| ----------------------- | ---------------------------------------------------------- | -------------- | -------- |
| WEB_HOST                | バックエンドのWEBサーバのホスト名                          | 空白           |
| WEB_PORT                | バックエンドのWEBサーバのポート                            | 8080           |
| WEB_ALLOW_ORIGIN        | WEBサーバのヘッダー情報のAccess-Control-Allow-Originの設定 | \*             |
| DB_TYPE                 | (設定変更不可)SQLサーバの種類指定                          | sqlite         |
| DB_FILE                 | SQLiteの保存ファイル名                                     | test.db        |
| JWT_TOKEN_HOUR_LIFESPAN | JWTトークンの有効期間                                      | 24             | 時間表示 |
| JWT_API_SECRET          | JWTトークンのシークレットキー                              | auth-go-sample |
| AUTH_PEPPER             | パスワードを暗号化する際のペッパー文字列                   | auth-go-sample |

## 各APIについて

以下のWeb APIを用意している
|API|Method|説明|オプション|備考|
|--|--|--|--|--|
|/|GET|
|/|OPTIONS|CORS対応のオプション
|/app/v1/login|GET|ヘッダーのトークン情報から
|/app/v1/login|POST|受け取ったJSONデータからユーザ認証を実施してトークン情報を返す|
|/app/v1/logout|POST|受け取ったJSONデータからユーザのログアウト処理を実施|
|/app/v1/user|PUT|受け取ったJSONデータからユーザ登録を実施|adminのみ
|/app/v1/user|POST|受け取ったJSONデータからユーザ登録を更新|admin|未実装|
|/app/v1/user/[:id]GET|id情報をもとにしたユーザ情報の取得|adminと特定ユーザのみ
|/app/v1/user/[:id]POST|id情報をもとにユーザ情報を更新|admin用
|/app/v1/user/[:id]|DELETE|id情報をもとにユーザ情報を削除|admin用で自分以外を削除
|/app/v1/user/list|GET|ユーザのリストを取得|admin用

以下のWebページの構成は以下の通りとなる
|API|説明|備考|
|--|--|--|
|/|ベースのインデックスページ||
|/login|ログイン用のテストページ||
|/home|1userとadmin用の表示テストページ||
|/page|gestとuser、admin用の表示テストページ||
|/edit|admin用のユーザ管理ページ||
