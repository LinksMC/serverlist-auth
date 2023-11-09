package main

import (
	"log/slog"

	"github.com/LinksMC/serverlist-auth/data"
	"github.com/sandertv/gophertunnel/minecraft"
)

func main() {
	slog.Info("サーバーを起動します...")
	listener, err := minecraft.ListenConfig{}.Listen("raknet", "0.0.0.0:19132")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(c.(*minecraft.Conn), listener)
	}
}

// クライアントの接続を処理
func handleConn(conn *minecraft.Conn, listener *minecraft.Listener) {
	// 接続情報取得
	indetity := conn.IdentityData()
	clientData := conn.ClientData()
	slog.Info("クライアントが接続しました", "Name", indetity.DisplayName, "XUID", indetity.XUID, "OS", data.GetDeviceOSName(clientData.DeviceOS))
	// TODO: DB操作
	// クライアントの接続を切断
	listener.Disconnect(conn, "connection lost")
}
