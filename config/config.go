package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type ConfigMap struct {
	DBConfig         PostgreSQLConfig
	MailerConfig     SendInBlueConfig
	MigrationFileURL string
	CacheConfig      InmemCacheConfig
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

type InmemCacheConfig struct {
	TTLInSecond       time.Duration
	PurgeTimeInSecond time.Duration
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

	fmt.Println(config)

	return &config, nil
}
