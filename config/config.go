package config

import "os"

var (
	ServerAddress      = os.Getenv("SERVER_ADDRESS")
	DBConnectionString = os.Getenv("DB_CONN")
	ImagesFilePath     = os.Getenv("IMAGES_FILEPATH")
	CognitiveURL       = os.Getenv("COGNITIVE_URL")
	CognitiveKey       = os.Getenv("COGNITIVE_KEY")
	SuccessPage        = os.Getenv("SUCCESS_PAGE")
)
