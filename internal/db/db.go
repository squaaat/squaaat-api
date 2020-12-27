package db

import (
	"database/sql"
	"fmt"
	gormLogger "gorm.io/gorm/logger"
	"os"
	baseLog "log"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/squaaat/squaaat-api/internal/config"
)

type Client struct {
	AppConfig *config.AppConfig
	Config    *config.ServiceDBConfig
	DB        *gorm.DB
}

func New(cfg *config.ServiceDBConfig, appcfg *config.AppConfig) (*Client, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)"+
			"/%s"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
	)

	var defaultLogger gormLogger.Interface
	if appcfg.Env == "alpha" {
		defaultLogger = gormLogger.New(
			baseLog.New(os.Stdout, "\r\n", baseLog.LstdFlags),
			gormLogger.Config{
				SlowThreshold: time.Second,   // Slow SQL threshold
				LogLevel:      gormLogger.Info, // Log level
				Colorful:      true,         // Disable color
			},
		)
	} else {
		defaultLogger = gormLogger.Default
	}

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:        dsn,
			DriverName: cfg.Dialect,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "sq_", // table name prefix, table for `User` would be `t_users`
				SingularTable: true,  // use singular table name, table for `User` would be `user` with this option enabled
			},
			Logger: defaultLogger,
		},
	)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second)

	return &Client{
		DB:        db,
		Config:    cfg,
		AppConfig: appcfg,
	}, nil
}

func Initialize(cfg *config.ServiceDBConfig) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	db, err := sql.Open(cfg.Dialect, dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", cfg.Schema))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Clean() {
	if c.AppConfig.Env != "alpha" {
		err := errors.New("Clean command only accept 'alpha' env")
		log.Fatal().Err(err).Send()
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/"+
			"?charset=utf8mb4&parseTime=True&loc=Local",
		c.Config.Username,
		c.Config.Password,
		c.Config.Host,
		c.Config.Port,
	)

	db, err := sql.Open(c.Config.Dialect, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s;", c.Config.Schema))
	if err != nil {
		panic(err)
	}

}
