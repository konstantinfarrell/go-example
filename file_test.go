package gox_test

import (
	"fmt"
	"testing"
	"reflect"
	
	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/helpers"
)

func TestToJson(t *testing.T){
	file := helpers.ExampleFile()
	expected := jsonResponse(file)
	actual, _ := file.ToJson()

	if actual != expected {
		t.Log("Actual:" + actual + "\nExpected:" , expected)
		t.Fail()
	}
}

func TestFromJson(t *testing.T){
	var f gox.File
	expected := helpers.ExampleFile()
	actual, _ := f.FromJson(jsonResponse(expected))

	if !reflect.DeepEqual(actual, expected) {
		err := fmt.Sprintf("Actual: %+v\nExpected: %+v", actual, expected)
		t.Log(err)
		t.Fail()
	}
}

func jsonResponse(f *gox.File) string {
	// TODO: figure out why Json.Marshal seems to create a different value for []byte()
	template := `{"FileId":"%s","Filename":"%s","Path":"%s","Permissions":"%s","Created":"%s","Modified":"%s","Data":"ZmlsZSBjb250ZW50cw==","Received":"%s"}`
	json := fmt.Sprintf(template, f.FileId, f.Filename, f.Path, f.Permissions, f.Created, f.Modified, f.Received)
	return json
}