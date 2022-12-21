package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Logger Global representation of log
var Logger = NewConsole(true)

// NewConsole initialize the global configuration of log
func NewConsole(isDebug bool) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("[%s]", 1)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s", 1)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return zerolog.New(output).With().Timestamp().Logger()
}
