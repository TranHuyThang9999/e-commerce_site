package configs

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configs struct {
	UrlImgbb      string `json:"url_imgbb"`
	DataSource    string `json:"data_source"`
	Port          string `json:"port"`
	AccessSecret  string `json:"access_secret"`
	ExpireAccess  string `json:"expire_access"`
	RefreshSecret string `json:"refresh_secret"`
	ExpireRefresh string `json:"expire_refresh"`
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
