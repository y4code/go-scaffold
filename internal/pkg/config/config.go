package config

import (
	_ "embed"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/y4code/go-scaffold/pkg/reflect"
)

//go:embed config.toml
var configFile string

type Option func(*Config)

type Config struct {
	ServiceName                string
	TomlTag                    string
	MaxIdle                    int
	MaxOpen                    int
	Profile                    string `json:"profile"`
	DBHost                     string `json:"db.host"`
	DBName                     string `json:"db.name"`
	DBUsername                 string `json:"db.username"`
	DBPassword                 string `json:"db.password"`
	ConfigCenterServiceHost    string `json:"configcenter-service.host"`
	AdminServiceHost           string `json:"admin-service.host"`
	AdminServiceGetServiceCard string `json:"admin-service.get-service-card"`
	UrpServiceHost             string `json:"urp-service.host"`
}

func NewConfig(serviceName string) *Config {
	return &Config{
		ServiceName: serviceName,
		TomlTag:     "json",
		MaxIdle:     5,
		MaxOpen:     100,
	}
}

func (c *Config) LoadLocalConf() *Config {
	tree, _ := toml.Load(configFile)
	assignment := reflect.NewReflectOperator(c, tree, c.TomlTag)
	assignment.AssignValueByTagFromToml()
	return c
}

func (c *Config) OverrideWithRemoteConf() *Config {
	loadDataFromConfigCenter(c.ConfigCenterServiceHost, c.Profile, c.ServiceName, &c)
	return c
}

func (c *Config) PrintAllConf() {
	fmt.Printf("%+v", c)
}
