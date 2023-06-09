# golang-rest-example

# my version

```
Docker version 20.10.12, build e91ed57
docker-compose version 1.29.2, build 5becea4c
```

# setup

## install

- go
  - https://go.dev/
- docker desktop
  - https://www.docker.com/products/docker-desktop/
- direnv(mac の場合)
  - https://qiita.com/toshichanapp/items/c5d76144bf6938207f4e

## run DB/DB コンテナを起動

```sh
$ docker-compose up -d #DBコンテナ起動
$ docker-compose down #DBコンテナ停止
```

## run app/api server を起動

1. setup environment/環境変数を設定(内容は.env.sample ファイル参照)
   ```sh
   $ export PORT=<PORT>
   $ export MYSQL_DSN=<MYSQL_DSN>
   $ export ALLOW_CORS_ORIGIN=<ALLOW_CORS_ORIGIN>
   ```
2. command/ローカルで api server を起動
   ```sh
   $ make run #go run ./cmd/api/main.go
   ```
3. check endpoint/ブラウザで動作確認
   ```sh
   $ localhost:<PORT>/version #ex  localhost:8080/version
   $ localhost:<PORT>/healthz #ex localhost:8080/healthz
   ```

- 画面キャプチャ
  - /version
    - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D--version
  - /healthz
    - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D-healthz

# test with linux/go でテストを実行

```sh
$ make test #go test ./...
```

# curl/http request/HTTP リクエスト(動作確認結果のキャプチャあり)

```sh
$ curl -X GET localhost:8080/version #バージョン確認(動作確認も兼ねています)
$ curl -X GET localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserを取得
$ curl -X GET localhost:8080/v1/user/all -H "Content-Type: application/json" #userテーブルの一覧(全件)を取得
$ curl -X POST localhost:8080/v1/user -H "Content-Type: application/json" --data-raw '{"name": "user"}' #usersテーブルに指定したnameのuserレコードを追加
$ curl -X DELETE localhost:8080/v1/user/{id} -H "Content-Type: application/json" #idを指定して該当のuserをテーブルから削除

```

- 画面キャプチャ
  - https://github.com/taaaaakahiro/golang-rest-example/wiki/%5B%E5%8B%95%E4%BD%9C%E7%A2%BA%E8%AA%8D%5D%E7%94%BB%E9%9D%A2%E3%82%AD%E3%83%A3%E3%83%97%E3%83%81%E3%83%A3

# Architecture

## DDD/Clean Architecture

### pkg

- command: 各 pkg を初期化
- config: 設定を管理
- domain: 目的の明確化
- handler: ルーティング
- infrastructure
  - persistence: クエリ
- io: DB 接続
- middleware: リクエストの前処理等
- service: 複数の UseCase を纏める
- server: API Server の設定
- version: tool バージョン管理

# 追加予定の機能

- [x] テンプレートエンジン 2023/2/6 up
- [ ] csv インポート
- [ ] csv エクスポート
- [ ] ファイルアップロード
- [ ] メール送信
- [x] gorilla → chi 2023/2/11 up
