package config

import "time"

type Config struct {
	Server           ServerConfig      `yaml:"server"`
	Database         PostgresConfig    `yaml:"database"`
	Mailer           SendInBlueConfig  `yaml:"mailer"`
	MigrationFileUrl string            `yaml:"migrationFileUrl"`
	Cache            CacheConfig       `yaml:"cache"`
	FeatureFlag      FeatureFlagConfig `yaml:"featureFlag"`
}

type ServerConfig struct {
	Name                 string `yaml:"name"`
	Port                 string `yaml:"port"`
	ShutdownTimeoutInSec int64  `yaml:"shutdownTimeoutInSec"`
}

type SendInBlueConfig struct {
	ApiKey     string `yaml:"apiKey"`
	PartnerKey string `yaml:"partnerKey"`
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Ssl      string `yaml:"ssl"`
}
type CacheConfig struct {
	Inmem InmemCacheConfig `yaml:"inmem"`
}

type InmemCacheConfig struct {
	TtlInSecond       time.Duration `yaml:"ttlInSecond"`
	PurgeTimeInSecond time.Duration `yaml:"purgeTimeInSecond"`
}

type FeatureFlagConfig struct {
	RunMigrations bool `yaml:"runMigrations"`
}
