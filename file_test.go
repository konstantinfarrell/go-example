package gox_test

import (
	"fmt"
	"testing"
	"time"
	"reflect"
	
	"github.com/google/uuid"

	"github.com/konstantinfarrell/go-example"
)

func TestToJson(t *testing.T){
	file, expected := exampleFile()

	actual, _ := file.ToJson()

	if actual != expected {
		t.Log("Actual:" + actual + "\nExpected:" , expected)
		t.Fail()
	}
}

func TestFromJson(t *testing.T){
	var f gox.File
	expected, json := exampleFile()
	actual, _ := f.FromJson(json)

	if !reflect.DeepEqual(actual, expected) {
		err := fmt.Sprintf("Actual: %+v\nExpected: %+v", actual, expected)
		t.Log(err)
		t.Fail()
	}
}

func exampleFile() (*gox.File, string) {
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

	// TODO: figure out why Json.Marshal seems to create a different value for []byte()
	template := `{"FileId":"%s","Filename":"%s","Path":"%s","Permissions":"%s","Created":"%s","Modified":"%s","Data":"ZmlsZSBjb250ZW50cw==","Received":"%s"}`
	json := fmt.Sprintf(template, fileId, filename, path, permissions, now, now, now)

	return model, json
}