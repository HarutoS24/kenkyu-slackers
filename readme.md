# 研究 slackers

## 環境

- Node.js v24.7.0
- Vue.js 3.5.18
- Go 1.24.2

### セットアップ手順

#### Node.js

- node nvm を使用してバージョンを揃える。
- nvm をインストール。
- frontend ディレクトリに移動し、nvm install を実行。

#### Vite

- backend ディレクトリに移動し、npm ci を実行。

#### Go

- go mod build

### 仕様

バックエンド api について

- get_feedback_from_GPT->レビューを返す。POST メソッドでアクセスする必要がある。JSON データを下のように投げるとアクセスできる。
  > ```json
  > {
  >   "text": "このたび弊社は、、",
  >   "industry_id": "9",
  >   "important_aspects": ["0", "1"]
  > }
  > ```
- 返却データの例
  > ```json
  > {
  >   "advice": "あなたのプレスリリースは以下の点で問題があります...",
  >   "improved_press": "#[業界初!!]..."
  > }
  > ```
