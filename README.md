# Departures Time
[Uber社のコーディングチャレンジ](https://github.com/uber-archive/coding-challenge-tools)を実装したもの

## 概要
バックエンドを中心に実装しました。
本来ならばフォーマットした公開データをDBに取り込み独自にデータを持つのがベストですが、時間がないため公開APIから駅データを取得するようにしています。
経路探索結果のデータそのものを提供してくれているAPIは有料であったたため、外部サービスの探索結果のURLを返すようにしています。

## 要件
- Node.js >= v16.14.2
- Docker Desktop >= 4.9.1

## 開発環境立ち上げ
### API
- `docker-compose.yml`の次の環境変数書き換えを駅すぱあとWebサービスで発行したAPIキーに書き換え`EKISPERT_API_KEY=<your api key>`
- VS Code以外で行う場合
  - `docker compose up --build`でコンテナを立ち上げ、`docker exec -it <container name> go run .`でAPIを起動する
- VS Codeを使用している場合
  - `cd departures-time-api && code .`
  - VS Codeが立ち上がったら[Reopen in Container]をクリックし、コンテナ内でVS Codeが立ち上がるのを待つ
  - コンテナ内のVS Codeが立ち上がるとGo拡張機能に必要なライブラリのインストールを促されるので[Install All]をクリック
  - `go run .`でAPIを起動する

### Front
- `cd departures-time-front`
- `npm install`
- 必要に応じて`./departures-time-front/src/api/base.ts`内の`BASE_PATH`を書き換え
  - ローカル `http://localhost:8000`
  - 本番 `https://departures-time-api.ohagi.link`
- `npm start`

## テスト
### API
- `go test ./...`
  - 公開APIのキーなどが設定されていない場合は`external`以下のテストは失敗します external以外のテストを見たい場合は一旦`external/*_test.go`を削除してから試してください
- カバレッジ出力
  - `go test ./... -coverprofile cover.out && go tool cover -html=cover.out`

## Frount用のAPIクライアントを生成する
- [openapi-generator](https://github.com/OpenAPITools/openapi-generator)をnpmでインストールする
- `cd departures-time-front`
- `openapi-generator-cli generate -g typescript-axios -i ../reference/departures-time-open-api.yaml -o ./src/api`
   - 一部エラーが出るので手動で直します

#
