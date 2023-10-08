package db

import "go.etcd.io/bbolt"

type App struct {
	DB *bbolt.DB
}

func InitDB() (*bbolt.DB, error) {
	db, err := bbolt.Open("url-shortner.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
