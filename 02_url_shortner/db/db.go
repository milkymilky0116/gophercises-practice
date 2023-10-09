package db

import "go.etcd.io/bbolt"

type App struct {
	DB *bbolt.DB
}

func InitDB(dbname string) (*bbolt.DB, error) {
	db, err := bbolt.Open(dbname, 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}
