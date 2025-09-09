# 研究slackers
## 環境
* Node.js v24.7.0
* Vue.js 3.5.18
* Go 1.24.2
###セットアップ手順
Node.js
node nvmを使用してバージョンを揃える。
・nvmをインストール。
・frontendディレクトリに移動し、nvm installを実行。
Vite
・backendディレクトリに移動し、npm ciを実行。
Go
・go mod build
###仕様
バックエンドapiについて
・check->POSTメソッドでアクセスする必要がある。JSONデータを下のように投げるとアクセスできる。
>```go
> {"text": "このたび弊社は、、","industry_id": "9", "important_aspects": ["0", "1"]}
>
・（未実装）アドバイスと修正例を返却するapi->POSTメソッドを使い、上述のJSONデータを投げると下のようなjsonデータを返却する。
>```go
>{"advice": "あなたのプレスリリースは以下の点で問題があります...", "improved_press": "#[業界初!!]..."}
>
