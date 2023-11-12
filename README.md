# serverlist-auth
LinksMCのMinecraftアカウント認証サーバー用アプリケーションです。  
Minecraft(統合版)からこのアプリケーションで立ち上げたサーバーに接続することで、Minecraftアカウントの認証リクエストを作成します。

## 使用法
一度サーバーを立ち上げると設定ファイル(config.json)が立ち上がります。
サーバーを停止して設定ファイルを編集した後、もう一度サーバーを立上てください。
```json:config.json
{
 "minecraft": {
  "address": "0.0.0.0:19132", // バインドするアドレス
  "motd": "[LinksMC]認証サーバー", // Minecraftで表示させるサーバー名
  "max_players": 100, // サーバーの最大待ち受け人数
  "message": "以下のコードを入力してください!\n[TOKEN]" // ユーザーに表示するテキスト
 },
 "internal": {
  "token_length": 8, // 認証コードの長さ
  "cache_time": 60, // 認証コードをキャッシュする時間
  "cache_size": 100 // キャッシュのサイズ
 }
}
```

また、.envファイルを作成し以下の環境変数を設定してください。
```env:.env
DATABASE_URL="データベースのURL"
PRISMA_CLIENT="go run github.com/steebchen/prisma-client-go" // このまま使用
```
