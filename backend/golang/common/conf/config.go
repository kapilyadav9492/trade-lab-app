package conf

import (
	"os"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port          string `yaml:"port"`
		Host          string `yaml:"host"`
		ServerAddress string `yaml:"serverAddress"`
	} `yaml:"server"`
}

var Cfg Config

func init() {
	f, err := os.Open("backend/golang/common/conf/config.yml")
	if err != nil {
		log.Fatal("2--->", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Cfg)
	if err != nil {
		log.Fatal("1", err)
	}
}
