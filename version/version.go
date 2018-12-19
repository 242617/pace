package version

import "os"

var (
	Application = os.Getenv("APPLICATION")
)
