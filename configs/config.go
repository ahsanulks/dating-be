package configs

import (
	_ "embed"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

//go:embed openapi.yaml
var OpenAPI []byte

var conf *ApplicationConfig

type ApplicationConfig struct {
	Server   Server   `mapstructure:"server"`
	Postgres DBConfig `mapstructure:"postgres"`
}

type Server struct {
	HTTP    ServerConfig `mapstructure:"http"`
	GRPC    ServerConfig `mapstructure:"grpc"`
	OpenAPI ServerConfig `mapstructure:"open_api"`
}

type DBConfig struct {
	Hostname string `mapstructure:"hostname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
}

type ServerConfig struct {
	Addr    string `mapstructure:"addr"`
	Timeout int    `mapstructure:"timeout"`
}

var basepath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(b)
}

func NewConfig() *ApplicationConfig {
	once := new(sync.Once)
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(basepath)
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %s", err))
		}

		if err := viper.Unmarshal(&conf); err != nil {
			panic(fmt.Errorf("vailed to load config %s", err))
		}
	})
	return conf
}
