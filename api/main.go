package main

import (
	"BackendGo/adapters/repo"
	"BackendGo/api/routing"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	gormDB, err := initDb()
	if err != nil {
		return err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()
	r := repo.NewSqliteRepo(gormDB)
	s := routing.NewServer(r)
	mux := s.SetUpRoutes()
	return http.ListenAndServe(":8080", mux)
}
