package models

//Using Singletor pattern

import (
	"log"
	"sync"
	u "github.com/Go-Run-React/Catalog-Microservice-Backend/utils"
	"github.com/Go-Run-React/Catalog-Microservice-Backend/data/storage"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type Repository struct {
	DB *gorm.DB
	l *log.Logger
}

var repo  *Repository

func GetRepository() *Repository {

	config := &storage.Config{
		Host: "database-postgres.ckzevsu2tnze.ap-south-1.rds.amazonaws.com",
		Port: "5432",
		User: "postgres",
		Password: "Rishi_58235",
		SSLmode: "disable",
		DBName: "phase1",
	}

    if repo == nil {
        lock.Lock()
        defer lock.Unlock()

        if repo == nil {
			l := u.Dlog
			db, err := storage.NewConnection(config)
			if err != nil {
				log.Fatal("Database connection failed")
				return nil
			}
            repo = &Repository{db, l}
			return repo
        } else {
			return repo
        }

    } else {
        return repo
    }
}
