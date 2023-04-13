package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		File     string `yaml:"file" envconfig:"DB_FILE"`
		Username string `yaml:"sql-user" envconfig:"DB_USERNAME"`
		Password string `yaml:"sql-pass" envconfig:"DB_PASSWORD"`
		Database string `yaml:"sql-database" envconfig:"DB_DATABASE"`
		Server   string `yaml:"sql-server" envconfig:"DB_SERVER"`
		Port     int    `yaml:"sql-port" envconfig:"DB_PORT"`
	} `yaml:"database"`

	Redis struct {
		Server string `yaml:"server" envconfig:"REDIS_SERVER"`
	} `yaml:"redis"`

	SSO struct {
		ClientId        string `yaml:"clientId" envconfig:"SSO_CLIENTID"`
		ClientSecret    string `yaml:"clientSecret" envconfig:"SSO_CLIENTSECRET"`
		BaseAddress     string `yaml:"baseAddress" envconfig:"SSO_BASEADDRESS"`
		BackBaseAddress string `yaml:"backBaseAddress" envconfig:"SSO_BACKBASEADDRESS"`
	} `yaml:"sso"`

	S3 struct {
		AwsKey       string `yaml:"aws_key" envconfig:"S3_AWSKEY"`
		AwsSecret    string `yaml:"aws_secret" envconfig:"S3_AWSSECRET"`
		HostBase     string `yaml:"s3_host_base" envconfig:"S3_HOSTBASE"`
		Urlgenerator string `yaml:"urlgenerator" envconfig:"S3_URLGENERATOR"`
		BucketName   string `yaml:"bucket" envconfig:"S3_BUCKET"`
	} `yaml:"s3"`
}

func readConfig(cfg *Config) {
	readFile(cfg)
	readEnv(cfg)
	fmt.Printf("%+v", cfg)
}

func readFile(cfg *Config) {
	fileName := "config.yml"
	doReadFile(cfg, fileName)
	s := os.Getenv("RUNENVIRONMENT")
	if len(s) > 0 {
		fileName = "config" + s + ".yml"
		doReadFile(cfg, fileName)
	}

}

func doReadFile(cfg *Config, fileName string) {
	f, _ := os.Open(fileName)
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	decoder.Decode(cfg)
}

func readEnv(cfg *Config) {
	envconfig.Process("", cfg)
}
