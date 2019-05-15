package web

import (
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var (
	configFileName  = "../conf/config.yaml"
	FrameworkConfig = FrameworkConf{}
)

type DebugConf struct {
	Enable bool	 `yaml:"Enable"`
	Port   int   `yaml:"Port"`
}

type ServerConf struct {
	IP            string `yaml:"IP"`
	Port          int    `yaml:"Port"`
	ReadTimeout   int    `yaml:"ReadTimeout"`
	WriteTimeout  int    `yaml:"WriteTimeout"`
	IdleTimeout   int    `yaml:"IdleTimeout"`
	MaxHeaderSize int    `yaml:"MaxHeaderSize"`
}

type LoggerConf struct {
	Writer     string `yaml:"Writer"`
	Level      string `yaml:"Level"`
	Format     string `yaml:"Format"`
	Rotate     bool   `yaml:"Rotate"`
	RotateType string `yaml:"RotateType"`
	LogPath    string `yaml:"LogPath"`
}

type LimiterConf struct {
	InterfaceName       string `yaml:"InterfaceName"`
	EnableRateLimit     bool   `yaml:"EnableRateLimit"`
	MaxRequestPerSecond int    `yaml:"MaxRequestPerSecond"`
}

type FrameworkConf struct {
	Application string        `yaml:"Application"`
	Debug       DebugConf     `yaml:"Debug"`
	Server      ServerConf    `yaml:"Server"`
	Logger      LoggerConf    `yaml:"Logger"`
	Limiter     []LimiterConf `yaml:"Limiter"`
}

func InitConfig() {
	file, err := os.OpenFile(configFileName, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &FrameworkConfig)
	if err != nil {
		panic(err)
	}
}
