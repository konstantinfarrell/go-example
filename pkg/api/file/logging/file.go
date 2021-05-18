package file

import (
	"time"

	"github.com/labstack/echo"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/api/file"
)

func New(svc file.Service, logger gox.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger: logger,
	}
}

type LogService struct {
	file.Service
	logger gox.Logger
}

const name = "file"

func (ls *LogService) Create(c echo.Context, request *gox.File) (response *gox.File, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Create file request", err,
			map[string]interface{}{
				"request":	request,
				"response": response,
				"duration": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Create(c, request)
}

func (ls *LogService) ReadAll(c echo.Context) (response *[]gox.File, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "List files request", err,
			map[string]interface{}{
				"response": response,
				"duration": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ReadAll(c)
}

func (ls *LogService) Read(c echo.Context, request *gox.File) (response *gox.File, err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Read file request", err,
			map[string]interface{}{
				"request":	request,
				"response": response,
				"duration": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Read(c, request)
}

func (ls *LogService) Delete(c echo.Context, request *gox.File) (err error) {
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			name, "Delete file request", err,
			map[string]interface{}{
				"request":	request,
				"duration": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Delete(c, request)
}