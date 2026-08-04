package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server     server     `yaml:"server"`
	Db         db         `yaml:"db"`
	Redis      redis      `yaml:"redis"`
	ImgSetting imgSetting `yaml:"imgSetting"`
	Log        log        `yaml:"log"`
}

type server struct {
	Address string `yaml:"address"`
	Mode    string `yaml:"mode"`
}

type db struct {
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type redis struct {
	address  string `yaml:"address"`
	password string `yaml:"password"`
}

type imgSetting struct {
	uploadDir string `yaml:"uploadDir"`
	ImgHost   string `yaml:"imgHost"`
}

type log struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
	Mode string `yaml:"mode"`
}

var Config *config

func init() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		panic(err)
	}
}
