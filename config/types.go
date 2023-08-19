package config

import "time"

type Config struct {
	Server           ServerConfig      `yaml:"server"`
	Database         DatabaseConfig    `yaml:"database"`
	Mailer           MailerConfig      `yaml:"mailer"`
	MigrationFileUrl string            `yaml:"migrationFileUrl"`
	Cache            CacheConfig       `yaml:"cache"`
	FeatureFlag      FeatureFlagConfig `yaml:"featureFlag"`
	Jwt              JwtConfig         `yaml:"jwt"`
}

type ServerConfig struct {
	Name                 string `yaml:"name"`
	Port                 string `yaml:"port"`
	ShutdownTimeoutInSec int64  `yaml:"shutdownTimeoutInSec"`
	SecretKey            string `yaml:"secretKey"`
}

type MailerConfig struct {
	SendInBlue SendInBlueConfig `yaml:"sendInBlue"`
}

type SendInBlueConfig struct {
	ApiKey     string `yaml:"apiKey"`
	PartnerKey string `yaml:"partnerKey"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
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

type JwtConfig struct {
	AccessExpiryInHour  time.Duration `yaml:"accessExpiryInHour"`
	RefreshExpiryInHour time.Duration `yaml:"refreshExpiryInHour"`
}
