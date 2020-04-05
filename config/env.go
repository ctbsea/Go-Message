package config

import (
	"github.com/ctbsea/Go-Message/util"
	"github.com/kataras/iris"
	"log"
)

type RunConfig struct {
	IP   string `yaml:"host"`
	Port string `yaml:"port"`
}

type MySQLConfig struct {
	IP           string `yaml:"host,omitempty"`
	Port         string `yaml:"port,omitempty"`
	User         string `yaml:"user,omitempty"`
	Password     string `yaml:"pass,omitempty"`
	Database     string `yaml:"db,omitempty"`
	Debug        bool `yaml:"debug,omitempty"`
	MaxIdleConns int `yaml:"maxIdleConns,omitempty"`
	MaxOpenConns int `yaml:"maxOpenConns,omitempty"`
}

type PathConfig struct {
	Runtime string
}

type EnvConfig struct {
	Env string `yaml:"env"`
}

type GateWay struct {
	LimiterOneSec float64 `yaml:"limiterOneSec"`
}

type Config struct {
	Run   *RunConfig
	MySQL *MySQLConfig
	Path  *PathConfig
	GateWay *GateWay
	Env *EnvConfig
}

//tmpConfig point
func InitConfig(config map[string]interface{}, tag string, tmpConfig interface{}, callback func(interface{})) {
	configValue, ok := config[tag]
	if !ok {
		log.Fatal(tag + " Config No Exists")
	}
	util.YamlInterfaceToStruct(configValue, tmpConfig)
	if callback != nil {
		callback(tmpConfig)
	}
}

func checkRunConfig(config interface{}) {
	if config.(*RunConfig).IP == "" {
		config.(*RunConfig).IP = "127.0.0.1"
	}
	if config.(*RunConfig).Port == "" {
		config.(*RunConfig).Port = "8888"
	}
}

func  checkEnvConfig(config interface{}) {
	if config.(*EnvConfig).Env == "" {
		config.(*EnvConfig).Env = "dev"
	}
}


func InitAllConfig(app *iris.Application) Config {
	env := iris.YAML("./web/env.yml")
	otherConfig := env.GetOther()
	var config Config

	mysqlConfig := &MySQLConfig{}
	InitConfig(otherConfig, "mysql", mysqlConfig, nil)
	config.MySQL = mysqlConfig

	runConfig := &RunConfig{}
	InitConfig(otherConfig, "run", runConfig, checkRunConfig)
	config.Run = runConfig

	runtime := "./runtime/"
	config.Path = &PathConfig{Runtime: runtime}

	gateway := &GateWay{}
	InitConfig(otherConfig, "gateway", gateway, nil)
	config.GateWay = gateway

	envConfig := &EnvConfig{}
	InitConfig(otherConfig, "env", envConfig, checkEnvConfig)
	config.Env = envConfig
	return config
}


