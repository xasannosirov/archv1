package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	PostgresDB      string `yaml:"postgres_db"`
	PostgresUser    string `yaml:"postgres_user"`
	PostgresPWD     string `yaml:"postgres_pwd"`
	PostgresPort    string `yaml:"postgres_port"`
	PostgresHost    string `yaml:"postgres_host"`
	PostgresSSLMode string `yaml:"postgres_ssl_mode"`

	RedisHost string `yaml:"redis_host"`
	RedisPort string `yaml:"redis_port"`
	RedisDB   int    `yaml:"redis_db"`
	RedisPWD  string `yaml:"redis_pwd"`

	HttpHost   string `yaml:"http_host"`
	HttpPort   string `yaml:"http_port"`
	CtxTimeout string `yaml:"ctx_timeout"`
	GinMode    string `yaml:"gin_mode"`

	AuthConfigPath string `yaml:"auth_config_path"`
	CSVFilePath    string `yaml:"csv_file_path"`

	AccessTTL  string `yaml:"access_ttl"`
	RefreshTTL string `yaml:"refresh_ttl"`
	JWTSecret  string `yaml:"jwt_secret"`
}

func NewConfig() *Config {
	c := &Config{}
	yamlFile, err := os.ReadFile("./internal/pkg/config/config.yaml")
	if err != nil {
		log.Fatalf("yamlFile.ReadFile err %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
