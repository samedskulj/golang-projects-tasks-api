package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	sqlStorage := NewMySQLStorage(mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Net:                  Envs.DBAddress,
		DBName:               Envs.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPIServer(":3000", store)
	api.Serve()
}
