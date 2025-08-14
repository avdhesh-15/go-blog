package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DbPool *pgxpool.Pool

func DbConfig() {

	err := godotenv.Load("../.env")
	if err != nil {
		println("Error in db env: ", err)
	}

	DbPool, err = pgxpool.New(context.Background(), os.Getenv("DB_STRING"))

	if err != nil {
		log.Fatal("Error occured in the DB connection: ", err)
	}

}
