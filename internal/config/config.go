package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	PROJECT = "squaaat"
	APP     = "squaaat-api"
)

func MustInit(e string, cicd bool) *Config {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	s := ssm.New(sess)
	param, err := s.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(fmt.Sprintf("/%s/%s/%s/env", PROJECT, APP, e)),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	value := *(param.Parameter.Value)
	viper.SetConfigType("yaml")
	viper.ReadConfig(strings.NewReader(value))

	return &Config{
		Version:    viper.GetString("version"),
		CICD:       cicd,
		App:        newAppConfig(),
		ServerHTTP: newServerHTTPConfig(),
		ServiceDB:  newServiceDBConfig(),
	}
}

type Config struct {
	Version    string
	CICD       bool
	App        *AppConfig
	ServerHTTP *ServerHTTPConfig
	ServiceDB  *ServiceDBConfig
}

func newServerHTTPConfig() *ServerHTTPConfig {
	return &ServerHTTPConfig{
		Port:    viper.GetString("env.server_http.port"),
		Timeout: viper.GetDuration("env.server_http.timeout"),
	}
}

func newServiceDBConfig() *ServiceDBConfig {
	return &ServiceDBConfig{
		Host:     viper.GetString("env.service_db.host"),
		Port:     viper.GetString("env.service_db.port"),
		Dialect:  viper.GetString("env.service_db.dialect"),
		Schema:   viper.GetString("env.service_db.schema"),
		Username: viper.GetString("env.service_db.username"),
		Password: viper.GetString("env.service_db.password"),
	}
}

func newAppConfig() *AppConfig {
	pemKeyLines := strings.Split(viper.GetString("env.app.rsa_private_pem"), "\\n")
	for i, line := range pemKeyLines {
		pemKeyLines[i] = strings.TrimSpace(line)
	}

	return &AppConfig{
		Env:           viper.GetString("env.app.env"),
		Debug:         viper.GetBool("env.app.debug"),
		Project:       viper.GetString("env.app.project"),
		AppName:       viper.GetString("env.app.app_name"),
		RSAPrivatePem: strings.Join(pemKeyLines, "\n"),
	}
}

type ServerHTTPConfig struct {
	Port    string
	Timeout time.Duration
}

type ServiceDBConfig struct {
	Host     string
	Port     string
	Dialect  string
	Schema   string
	Username string
	Password string
}

type AppConfig struct {
	Env           string
	Debug         bool
	Project       string
	AppName       string
	RSAPrivatePem string
}
