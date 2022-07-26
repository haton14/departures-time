# 現在地の最寄駅を選択し、目的駅までの経路探索結果を生成する

## 技術について
### 使用したもの
#### 開発
- React v18 (業務で使用したことなくジュニアレベルです)
- TypeScript v4.6 (業務で多少書いたことがあるレベルです)
- [openapi-generator](https://github.com/OpenAPITools/openapi-generator) (API通信のクライアント生成に使用)
#### インフラ
- [Vercel](https://vercel.com/)

## 改善点
- CSSを使用する
- テストを書く
- 不要なライブラリを整備
- CI, CDを設定
- APIのパスを環境変数化
- 開発環境コンテナ化して環境構築を容易にするべき
- ロギングの実装
