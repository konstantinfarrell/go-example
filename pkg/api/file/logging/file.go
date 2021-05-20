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

func (ls *LogService) Create(fc chan gox.FileChannel, c echo.Context, request *gox.File) {
	var response *[]gox.File
	var err error
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
	go ls.Service.Create(fc, c, request)
	result := <- fc
	response = result.File
	err = result.Err
	return
}

func (ls *LogService) ReadAll(fc chan gox.FileChannel, c echo.Context) {
	var response *[]gox.File
	var err error
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
	go ls.Service.ReadAll(fc, c)
	result := <- fc
	response = result.File
	err = result.Err
}

func (ls *LogService) Read(fc chan gox.FileChannel, c echo.Context, request *gox.File) {
	var response *[]gox.File
	var err error
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
	go ls.Service.Read(fc, c, request)
	result := <- fc
	response = result.File
	err = result.Err
}

func (ls *LogService) Delete(fc chan gox.FileChannel, c echo.Context, request *gox.File) {
	var err error
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
	go ls.Service.Delete(fc, c, request)
	result := <- fc
	err = result.Err
}