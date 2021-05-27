package mock

import (
	"github.com/konstantinfarrell/go-example"
)


type File struct{}

func (f File) Create(kc *interface{}, streamName string, file *gox.File, partitionKey string) (string, error){
	return "Ok", nil
}

func (f File) Delete(kc *interface{}, streamName string, string FileId, partitionKey string) (string, error){
	return "Ok", nil
}