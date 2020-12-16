package db

import (
	"database/sql"

	"github.com/ferdoran/go-sro-framework/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnAccount() (db *sql.DB) {
	db, errDb := sql.Open(config.GlobalConfig.DB.Account.Driver, config.ConnStringAccount())
	if errDb != nil {
		panic(errDb.Error())
	}
	return db
}

func OpenConnShard() (db *sql.DB) {
	db, errDb := sql.Open(config.GlobalConfig.DB.Shard.Driver, config.ConnStringShard())
	if errDb != nil {
		panic(errDb.Error())
	}
	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
