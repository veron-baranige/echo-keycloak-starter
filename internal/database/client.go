package db

import (
	"database/sql"
	"fmt"

	"github.com/veron-baranige/echo-keycloak-starter/config"
)

var (
	Client *Queries
	conn   *sql.DB
)

func SetupClient() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Env.DbUser,
		config.Env.DbPassword,
		config.Env.DbHost,
		config.Env.DbPort,
		config.Env.DbName,
	)

	dbConn, err := sql.Open(config.Env.DbDriver, connStr)
	if err != nil {
		return err
	}

	Client = New(dbConn)
	conn = dbConn

	return nil
}

func GetDbConn() *sql.DB {
	return conn
}