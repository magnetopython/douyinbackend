package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Mysql Mysql `yaml:"mysql"`
}

type Mysql struct {
	UserName  string `yaml:"username"`
	Password  string `yaml:"password"`
	IpAddress string `yaml:"ipaddress"`
	Port      string `yaml:"port"`
	DbName    string `yaml:"dbname"`
}

var C Config

func ConfigInit() error {
	dataBytes, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("读取yaml文件失败")
		return err
	}
	err = yaml.Unmarshal(dataBytes, &C)
	if err != nil {
		fmt.Println("解析yaml文件失败")
		return err
	}
	return nil
}
