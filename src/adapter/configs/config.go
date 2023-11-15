package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configs struct {
	UrlImgbb   string `json:"url_imgbb"`
	DataSource string `json:"data_source"`
	Port       string `json:"port"`
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}
