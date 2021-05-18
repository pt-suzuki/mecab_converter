package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pt-suzuki/mecab_converter/application/config"
	"log"
)

func getDbConnect() *sql.DB {
	conf := config.NewConfig()
	db, _ := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d converter=%s password=%s dbname=%s sslmode=disable",
			conf.DB.Host,conf.DB.Port,conf.DB.User,conf.DB.Password,conf.DB.Name))

	return db
}

func Query(query string) *sql.Rows{
	db := getDbConnect()
	defer db.Close()

	result, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func QueryTransactionExec(query string, param ...interface{}) sql.Result{
	db := getDbConnect()
	defer db.Close()


	result, err2 := db.Exec(query, param...)

	if err2 != nil {
		log.Fatal(err2)
	}

	return result
}


