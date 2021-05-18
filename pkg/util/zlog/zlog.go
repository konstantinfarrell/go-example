package zlog

import (
	"os"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

type Log struct {
	logger *zerolog.Logger
}

func New() *Log {
	z := zerolog.New(os.Stdout)
	return &Log{
		logger: &z,
	}
}

func (l *Log) Log(context echo.Context, source string, message string, err error, params map[string]interface{}) {
	if params == nil {
		params = make(map[string]interface{})
	}

	params["source"] = source
	if err != nil {
		params["error"] = err
		l.logger.Error().Fields(params).Msg(message)
		return
	}

	l.logger.Info().Fields(params).Msg(message)
}