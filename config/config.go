package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type chainConfig struct {
	Url string `json:"url"`
}

type Config struct {
	ActiveChains []string               `json:"active_chains"`
	ChainConfigs map[string]chainConfig `json:"chain_configs"`
}

type ConfToml struct {
	Production Config `toml:"production"`
	Staging    Config `toml:"staging"`
	Local      Config `toml:"local"`
}

func InitConfig() (Config, error) {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		log.Fatal("Fail to open config.json", err)
	}

	defer jsonFile.Close()
	var config Config
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)

	return config, nil
}
