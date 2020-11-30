package config

import (
	"os"
)

var env ENV

// Init ...
func Init() {
	env = ENV{}

	// Common
	env.Env = os.Getenv("ENV")

	// Port
	env.Port.App = os.Getenv("PORT")

	// Database
	env.Database.URI = os.Getenv("DATABASE_URI")
	env.Database.Name = os.Getenv("DATABASE_NAME")
	env.Database.Auth.Mechanism = os.Getenv("DATABASE_AUTH_MECHANISM")
	env.Database.Auth.Source = os.Getenv("DATABASE_AUTH_SOURCE")
	env.Database.Auth.Username = os.Getenv("DATABASE_AUTH_USERNAME")
	env.Database.Auth.Password = os.Getenv("DATABASE_AUTH_PASSWORD")

	// Redis
	env.Redis.URI = os.Getenv("REDIS_URI")
	env.Redis.Password = os.Getenv("REDIS_PASSWORD")

	// Auth
	env.Auth.SecretKey = os.Getenv("AUTH_SECRET_KEY")
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}

// IsTest ...
func IsTest() bool {
	return env.Env == "test"
}

// IsDev ...
func IsDev() bool {
	return env.Env == "dev"
}

// IsProduction ...
func IsProduction() bool {
	return env.Env == "production"
}
