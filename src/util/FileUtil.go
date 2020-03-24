package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 从文件中读取配置
func ReadConfigFromFile(result *[]ConnectInfo) interface{} {
	executePath, err := os.Executable()
	if IsEmpty(&executePath) {
		return "get execute path fail!"
	}
	configFile := filepath.Join(filepath.Dir(executePath), "./config.json")
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("读取配置文件的路径为:%s", configFile)
		return err
	}
	err = json.Unmarshal(file, result)
	if err != nil {
		return err
	}
	return nil
}
