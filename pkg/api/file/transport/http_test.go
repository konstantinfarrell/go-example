package transport_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/mocks"
	"github.com/konstantinfarrell/go-example/pkg/util/helpers"
	"github.com/konstantinfarrell/go-example/pkg/util/server"

	"github.com/konstantinfarrell/go-example/pkg/api/file/transport"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	// Arrange

	// Create server
	s := server.New()
	g := s.Group("")

	// Create example file, json, & file list for arguments and response
	f := helpers.ExampleFile()
	fs, err := f.ToJson()	
	flist := []gox.File{*f}

	// Mock service methods
	service := &mocks.Service{}
	service.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	service.On("GetResult", mock.Anything).Return(&gox.FileChannel{&flist, nil})
	
	// Start service
	transport.NewHTTP(service, g)
	ts := httptest.NewServer(s)
	defer ts.Close()
	
	// Act

	// Send request to server 
	path := ts.URL + "/files"
	res, err := http.Post(path, "application/json", bytes.NewBufferString(fs))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Serialize response for comparison
	fileObject := new([]gox.File)
	if err := json.NewDecoder(res.Body).Decode(fileObject); err != nil {
		t.Fatal(err)
	}
	response, err := (*fileObject)[0].ToJson()

	// Assert
	assert.Equal(fs, response)
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)

	s := server.New()
	g := s.Group("")

	f1 := helpers.ExampleFile()
	fs1, err := f1.ToJson()
	f2 := helpers.ExampleFile()
	fs2, err := f2.ToJson()
	flist := []gox.File{*f1, *f2}

	service := &mocks.Service{}
	service.On("ReadAll", mock.Anything, mock.Anything).Return(nil)
	service.On("GetResult", mock.Anything).Return(&gox.FileChannel{&flist, nil})

	transport.NewHTTP(service, g)
	ts := httptest.NewServer(s)
	defer ts.Close()

	path := ts.URL + "/files"
	res, err := http.Get(path)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	fileObject := new([]gox.File)
	if err := json.NewDecoder(res.Body).Decode(fileObject); err != nil {
		t.Fatal(err)
	}
	resp1, err := (*fileObject)[0].ToJson()
	resp2, err := (*fileObject)[1].ToJson()

	assert.Equal(fs1, resp1)
	assert.Equal(fs2, resp2)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)

	s := server.New()
	g := s.Group("")

	f := helpers.ExampleFile()
	fs, err := f.ToJson()
	flist := []gox.File{*f}
	fId := &gox.File{FileId: f.FileId}

	service := &mocks.Service{}
	service.On("Read", mock.Anything, mock.Anything, fId).Return(nil)
	service.On("GetResult", mock.Anything).Return(&gox.FileChannel{&flist, nil})

	transport.NewHTTP(service, g)
	ts := httptest.NewServer(s)
	defer ts.Close()

	path := ts.URL + "/files/" + f.FileId
	res, err := http.Get(path)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	fileObject := new([]gox.File)
	if err := json.NewDecoder(res.Body).Decode(fileObject); err != nil {
		t.Fatal(err)
	}
	resp, err := (*fileObject)[0].ToJson()

	assert.Equal(fs, resp)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	s := server.New()
	g := s.Group("")

	f := helpers.ExampleFile()
	fId := &gox.File{FileId: f.FileId}

	service := &mocks.Service{}
	service.On("Delete", mock.Anything, mock.Anything, fId).Return(nil)
	service.On("GetResult", mock.Anything).Return(&gox.FileChannel{nil, nil})

	transport.NewHTTP(service, g)
	ts := httptest.NewServer(s)
	defer ts.Close()
	
	path := ts.URL + "/files/" + f.FileId
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", path, nil)
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(http.StatusOK, res.StatusCode)
}