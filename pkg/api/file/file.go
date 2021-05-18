package file

import (
	"github.com/labstack/echo"

	"github.com/konstantinfarrell/go-example"
)


// Creates and deletes are sent to kinesis
func (f File) Create(c echo.Context, request *gox.File) (*gox.File, error) {
	_, err := f.aws.Create(f.kinesis, f.streamName, request, f.partitionKey)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (f File) Delete(c echo.Context, request *gox.File) error {
	_, err := f.aws.Delete(f.kinesis, f.streamName, request.FileId, f.partitionKey)
	if err != nil {
		return err
	}
	return nil
}

// Reads are sent to the DB
func (f File) Read(c echo.Context, request *gox.File) (*gox.File, error) {
	result, err := f.database.ReadFile(f.pgs, request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (f File) ReadAll(c echo.Context) (*[]gox.File, error) {
	result, err := f.database.ReadAllFiles(f.pgs)
	if err != nil {
		return nil, err
	}
	return result, nil
}