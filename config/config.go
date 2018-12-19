package config

import "os"

var (
	ServerAddress      = os.Getenv("SERVER_ADDRESS")
	DBConnectionString = os.Getenv("DB_CONN")
)
