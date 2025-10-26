package context

import (
	"api-requester/db"
	"database/sql"
	"log"
	"os"
)

// GLOBAL CONTEXT REFERENCE
type AppContext struct {
	DB     *sql.DB
	Logger *log.Logger
}

func NewAppContext() (*AppContext, error) {
	database, err := db.Connect()
	if err != nil {
		return nil, err
	}

	err = db.InitSchema(database)
	if err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	logger := log.New(logFile, "[APP] ", log.Ldate|log.Ltime|log.Lshortfile)

	return &AppContext{
		DB:     database,
		Logger: logger,
	}, nil
}
