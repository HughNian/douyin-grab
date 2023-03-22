package logger

import (
	"fmt"
	"os"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", os.Getenv("RUNTIME_ROOT_PATH"), os.Getenv("LOG_SAVE_PATH"))
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		os.Getenv("LOG_SAVE_NAME"),
		time.Now().Format(os.Getenv("TIME_FORMAT")),
		os.Getenv("LOG_FILE_EXT"),
	)
}
