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
	_"github.com/konstantinfarrell/go-example/pkg/util/postgres"

	_"github.com/konstantinfarrell/go-example/pkg/api/file"
	"github.com/konstantinfarrell/go-example/pkg/api/file/transport"
	//_ faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	//_ ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"

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
	assert.Equal(response, fs)
}