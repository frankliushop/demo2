package dataconfig

import (
	"encoding/json"
	"flag"

	"github.com/gobuffalo/packr"
)

//全域設定檔
var GlobalConfigData ConfigData

//資源檔
var GlobalBox packr.Box

//設定檔結構
type ConfigData struct {
	WorkType      string `json:"workType"`
	SqlConnection string `json:"sqlConnection"`
	Port          int    `json:"port"`
}

//初始化套件
func init() {
	var workType string
	flag.StringVar(&workType, "w", "dev", "work type")
	flag.Parse()
	GlobalBox := packr.NewBox("../configfile")
	fileName := "config_" + workType + ".json"
	configFile, _ := GlobalBox.FindString(fileName)
	_ = json.Unmarshal([]byte(configFile), &GlobalConfigData)
}
