package log

import (
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

func IsDebug(lvl string) bool {
	return strings.ToUpper(lvl) == "DEBUG"
}

func GetLogLevel(lvl string) log.Lvl {
	level := strings.ToUpper(lvl)
	switch level {
	case "DEBUG":
		return log.DEBUG
	case "INFO":
		return log.INFO
	case "WARN":
		return log.WARN
	case "ERROR":
		return log.ERROR
	case "OFF":
		return log.OFF
	}
	fmt.Errorf("Unable to parse log level string, %s", level)
	return 0
}