package config

import (
	"strings"

	"github.com/spf13/viper"
)

type ConfigMap struct {
	DBConfig     PostgreSQLConfig
	MailerConfig SendInBlueConfig
}

type SendInBlueConfig struct {
	ApiKey     string
	PartnerKey string
}

type PostgreSQLConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
	Sslmode  string
	Timezone string
}

func LoadConfig(path string) (*ConfigMap, error) {
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config ConfigMap
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
