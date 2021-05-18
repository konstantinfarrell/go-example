package gox

import "github.com/labstack/echo"

type Logger interface {
	Log(echo.Context, string, string, error, map[string]interface{})
}