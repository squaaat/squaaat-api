package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/squaaat/squaaat-api/internal/config"
)

type Client struct {
	Config *config.ServiceDBConfig
	DB *gorm.DB
}

func (c *Client) Initialize() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/" +
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

	_,err = db.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", c.Config.Schema))
	if err != nil {
		panic(err)
	}
}

func New(cfg *config.ServiceDBConfig) *Client {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)" +
			"/%s" +
			"?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Schema,
	)

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
			DriverName: cfg.Dialect,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: "sq_",   // table name prefix, table for `User` would be `t_users`
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			},
		},
	)
	if err != nil {
		fmt.Println()
		err = errors.WithStack(err)
		log.Fatal().Err(err).Send()
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Second)

	return &Client{
		DB: db,
		Config: cfg,
	}
}