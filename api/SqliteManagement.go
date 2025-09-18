package main

import (
	"BackendGo/adapters/repo"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file:database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&repo.DocumentRow{}); err != nil {
		return nil, err
	}
	return db, nil
}
