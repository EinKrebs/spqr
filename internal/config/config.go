package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	LogLevel string `default:"DEBUG"`// TODO usage
	Addr     string `required:"true"`
	ADMAddr  string `required:"true"`
	HttpAddr string `required:"true"`
	PROTO    string `required:"true"`
	SpqrData string `default:"docker/router/cfg.yaml"`
}

var (
	appConfig     AppConfig
	routingConfig RoutingConfig
)

func InitConfig(filePath string) error {
	if err := initAppConfig(); err != nil {
		return err
	}
	if err := initRoutingConfig(filePath); err != nil {
		return err
	}
	return nil
}

func GetAppConfig() *AppConfig {
	return &appConfig
}

func GetRoutingConfig() *RoutingConfig {
	return &routingConfig
}

func initAppConfig() error {
	err := envconfig.Process("spqr", &appConfig)
	if err != nil {
		return err
	}
	
	configBytes, err := json.MarshalIndent(appConfig, "", "  ")
	if err != nil {
		return err
	}
	log.Println("AppConfig:", string(configBytes))
	return nil
}

func initRoutingConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := yaml.NewDecoder(file).Decode(&routingConfig); err != nil {
		return err
	}

	configBytes, err := json.MarshalIndent(routingConfig, "", "  ")
	if err != nil {
		return err
	}
	log.Println("RoutingConfig:", string(configBytes))
	return nil
}
