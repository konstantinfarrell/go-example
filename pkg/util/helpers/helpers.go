package helpers

import (
	_ "fmt"
	_ "reflect"
	"time"

	"github.com/google/uuid"

	"github.com/konstantinfarrell/go-example"
)

func ExampleFile() *gox.File {
	fileId := uuid.New().String()
	filename := "filename"
	path := "C:/Users/foo/filename.txt"
	permissions := "rw"
	now := time.Now().UTC().Format(time.RFC3339)
	data := []byte("file contents")
	model := &gox.File{
		FileId: 		fileId,
		Filename:		filename,
		Path:			path,
		Permissions:	permissions,
		Created:		now,
		Modified:		now,
		Data:			data,
		Received:		now,
	}
	return model
}