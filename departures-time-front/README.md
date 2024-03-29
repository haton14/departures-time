# Departures Time Front
## 技術について
geolocationを使用して現在地を取得しています
### 使用したもの
#### 開発
- React v18 (業務で使用したことなくジュニアレベルです)
- TypeScript v4.6 (業務で多少書いたことがあるレベルです)
- [openapi-generator](https://github.com/OpenAPITools/openapi-generator) (API通信のクライアント生成に使用)
- @tsconfig/strictest/tsconfig.json (TypeScriptの型制約を厳しくするためのテンプレとして使用しています)
#### インフラ
- [Vercel](https://vercel.com/)

## 改善点
- 最寄駅一覧を取得する際に距離範囲を指定するようにすべきです (API側ではできるようにはしてあります)
- API通信呼び出しとロジックをもっと分割する
- CSSを使用する
- テストを書く
- 不要なライブラリを整備
- CI, CDを設定
- APIのパスを環境変数化
- 開発環境コンテナ化して環境構築を容易にするべき
- ロギングの実装
