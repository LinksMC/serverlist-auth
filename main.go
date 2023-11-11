package main

import (
	"context"
	"log/slog"

	"github.com/LinksMC/serverlist-auth/data"
	"github.com/LinksMC/serverlist-auth/gen"
	"github.com/LinksMC/serverlist-auth/prisma/db"
	"github.com/joho/godotenv"
	"github.com/sandertv/gophertunnel/minecraft"
)

func main() {
	// .envを読み込む
	loadEnv()

	// DB接続
	slog.Info("DBに接続します...")
	prisma := db.NewClient()
	if err := prisma.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := prisma.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
	slog.Info("DBに接続しました")

	// サーバー起動
	slog.Info("サーバーを起動します...")
	listener, err := getConfig().Listen("raknet", "0.0.0.0:19132")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	slog.Info("サーバーを起動しました")
	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(c.(*minecraft.Conn), listener, prisma)
	}
}

// クライアントの接続を処理
func handleConn(conn *minecraft.Conn, listener *minecraft.Listener, prisma *db.PrismaClient) {
	// 接続情報取得
	identity := conn.IdentityData()
	clientData := conn.ClientData()
	slog.Info("クライアントが接続しました", "Name", identity.DisplayName, "XUID", identity.XUID, "OS", data.GetDeviceOSName(clientData.DeviceOS), "IP", conn.RemoteAddr().String())

	// トークン保存 / 更新
	token := gen.CreateToken()
	request, err := prisma.MinecraftAuthRequest.UpsertOne(
		db.MinecraftAuthRequest.EditionMcid(
			db.MinecraftAuthRequest.Edition.Equals("bedrock"),
			db.MinecraftAuthRequest.Mcid.Equals(identity.XUID),
		),
	).Create(
		db.MinecraftAuthRequest.Edition.Set("bedrock"),
		db.MinecraftAuthRequest.Name.Set(identity.DisplayName),
		db.MinecraftAuthRequest.Mcid.Set(identity.XUID),
		db.MinecraftAuthRequest.Token.Set(token),
	).Update(
		db.MinecraftAuthRequest.Edition.Set("bedrock"),
		db.MinecraftAuthRequest.Name.Set(identity.DisplayName),
		db.MinecraftAuthRequest.Token.Set(token),
	).Exec(context.Background())
	if err != nil {
		slog.Error("トークンの保存 / 更新に失敗しました", "Error", err)
	}

	// クライアントの接続を切断
	listener.Disconnect(conn, "以下のコードを入力してください!\n"+request.Token)
}

// サーバー設定読み込み
func getConfig() minecraft.ListenConfig {
	config := minecraft.ListenConfig{
		AuthenticationDisabled: false,
		MaximumPlayers:         100,
		StatusProvider:         minecraft.NewStatusProvider("[LinksMC]認証サーバー"),
	}
	return config
}

// .envを読み込む
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("読み込み出来ませんでした: %v", err)
	}
}
