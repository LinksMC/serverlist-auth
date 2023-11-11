package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/creasty/defaults"
)

type MinecraftServerParams struct {
	Address    string `json:"address" default:"0.0.0.0:19132"`
	Motd       string `json:"motd" default:"[LinksMC]認証サーバー"`
	MaxPlayers int    `json:"max_players" default:"100"`
	Message    string `json:"message" default:"以下のコードを入力してください!\n[TOKEN]"`
}

type InternalServerParams struct {
	TokenLength int `json:"token_length" default:"8"`
}

type Config struct {
	Minecraft MinecraftServerParams `json:"minecraft"`
	Internal  InternalServerParams  `json:"internal"`
}

func CreateFile() {
	_config := &Config{}
	if err := defaults.Set(_config); err != nil {
		panic(err)
	}
	file, _ := json.MarshalIndent(_config, "", " ")
	_ = ioutil.WriteFile("settings.json", file, 0644)
}

func GetConfig() (Config, error) {
	zero := Config{}
	// デフォルト値を代入
	ini := Config{}
	if err := defaults.Set(&ini); err != nil {
		return zero, fmt.Errorf("初期値を代入できませんでした: %v", err)
	}
	// 設定ファイルを読み込む
	_, err := os.Open("config.json")
	if err != nil {
		decoded, err := json.MarshalIndent(ini, "", " ")
		if err != nil {
			return zero, fmt.Errorf("初期値をデコードできませんでした: %v", err)
		}
		if err := os.WriteFile("config.json", decoded, 0644); err != nil {
			return zero, fmt.Errorf("書き込みに失敗しました: %v", err)
		}
		return ini, nil
	}
	data, err := os.ReadFile("config.json")
	if err != nil {
		return zero, fmt.Errorf("設定ファイルを読み込めませんでした: %v", err)
	}
	var _config Config
	if err := json.Unmarshal(data, &_config); err != nil {
		return zero, fmt.Errorf("設定ファイルのデコードに失敗しました: %v", err)
	}
	return _config, nil
}
