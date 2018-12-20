package config

import "os"

var (
	ServerAddress      = os.Getenv("SERVER_ADDRESS")
	DBConnectionString = os.Getenv("DB_CONN")
	ImagesFilePath     = os.Getenv("IMAGES_FILEPATH")
	ImagesPath         = os.Getenv("IMAGES_PATH")
	CognitiveURL       = os.Getenv("COGNITIVE_URL")
	CognitiveKey       = os.Getenv("COGNITIVE_KEY")
)
