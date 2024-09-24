package config

import (
	"github.com/joho/godotenv"
	"os"
)

type (
	Container struct {
		App   *App
		DB    *DB
		HTTP  *HTTP
		MONGO *MONGO
	}

	App struct {
		Name string
		Env  string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
		Prefix         string
	}

	MONGO struct {
		MongoURL        string
		MongoDB         string
		MongoCollection string
	}
)

func New() (*Container, error) {
	godotenv.Load()

	return &Container{
		&App{
			Name: os.Getenv("APP_NAME"),
			Env:  os.Getenv("APP_ENV"),
		},
		&DB{
			Connection: os.Getenv("DB_CONNECTION"),
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			Name:       os.Getenv("DB_NAME"),
		},
		&HTTP{
			Env:    os.Getenv("HTTP_ENV"),
			URL:    os.Getenv("HTTP_URL"),
			Port:   os.Getenv("HTTP_PORT"),
			Prefix: os.Getenv("HTTP_PREFIX"),
		},
		&MONGO{
			MongoURL:        os.Getenv("MONGO_URL"),
			MongoDB:         os.Getenv("MONGO_DB"),
			MongoCollection: os.Getenv("MONGO_COLLECTION"),
		},
	}, nil
}
