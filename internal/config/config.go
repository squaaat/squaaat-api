package config

import (
	"fmt"
	"os"
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

func MustInit() {
	e := os.Getenv("SQ_ENV")

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

	Version = viper.GetString("version")
	App = newAppConfig()
	ServerHTTP = newServerHTTPConfig()
	Sentry = newSentryConfig()
	ServiceDB = newServiceDBConfig()
}

var (
	Version    string
	App        *AppConfig
	ServerHTTP *ServerHTTPConfig
	Sentry     *SentryConfig
	ServiceDB  *ServiceDBConfig
)

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

func newSentryConfig() *SentryConfig {
	return &SentryConfig{
		Enabled:      viper.GetBool("env.sentry.enabled"),
		DSN:          viper.GetString("env.sentry.dsn"),
		FlushTimeout: viper.GetDuration("env.sentry.flush_timeout"),
	}
}

func newAppConfig() *AppConfig {
	return &AppConfig{
		Env:     viper.GetString("env.app.env"),
		Debug:   viper.GetBool("env.app.debug"),
		Project: viper.GetString("env.app.project"),
		AppName: viper.GetString("env.app.app_name"),
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

type SentryConfig struct {
	Enabled      bool
	DSN          string
	FlushTimeout time.Duration
}

type AppConfig struct {
	Env     string
	Debug   bool
	Project string
	AppName string
}
