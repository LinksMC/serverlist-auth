package config

type MinecraftServerParams struct {
	Address string `json:"address" default:"0.0.0.0:19132"`
	Motd    string `json:"motd" default:"[LinksMC]認証サーバー"`
	Message string `json:"message" default:"以下のコードを入力してください!\n[TOKEN]"`
}

type InternalServerParams struct {
	TokenLength int `json:"token_length" default:"8"`
}

type Config struct {
	Minecraft MinecraftServerParams `json:"minecraft"`
	Internal  InternalServerParams  `json:"internal"`
}
