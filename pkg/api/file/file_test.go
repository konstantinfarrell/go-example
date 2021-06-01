package file_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/stretchr/testify/assert"
	"github.com/go-pg/pg/v9"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/mocks"
	"github.com/konstantinfarrell/go-example/pkg/api/file"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	"github.com/konstantinfarrell/go-example/pkg/util/helpers"
	faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

func exampleFile() *gox.File {
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

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	database := &postgres.Database{ Conn: &mocks.Databaser{}}

	sn := "foo"
	pk := "0"
	f := helpers.ExampleFile()
	data, _ := f.ToJson()
	formatted, _ := faws.FormatPayload(data, "create")
	payload := &kinesis.PutRecordInput {
		Data:			formatted,
		StreamName:		&sn,
		PartitionKey:	&pk,
	}
	ks := new(mocks.FileKinesis)
	ks.On("PutRecord", payload).Return(&kinesis.PutRecordOutput{}, nil)

	fileService := file.New(
		&mocks.Cacher{},
		database,
		&ffile.File{},
		ks,
		&faws.File{},
		sn,
		pk,
	)

	fc := make(chan gox.FileChannel)
	defer close(fc)
	go fileService.Create(fc, &mocks.Context{}, f)
	result := <- fc
	err := result.Err
	assert.Nil(err)
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	database := &postgres.Database{ Conn: &mocks.Databaser{}}

	sn := "foo"
	pk := "0"
	f := &gox.File{FileId: uuid.New().String()}
	data, _ := f.ToJson()
	formatted, _ := faws.FormatPayload(data, "delete")
	payload := &kinesis.PutRecordInput {
		Data:			formatted,
		StreamName:		&sn,
		PartitionKey:	&pk,
	}
	ks := new(mocks.FileKinesis)
	ks.On("PutRecord", payload).Return(&kinesis.PutRecordOutput{}, nil)

	fileService := file.New(
		&mocks.Cacher{},
		database,
		&ffile.File{},
		ks,
		&faws.File{},
		sn,
		pk,
	)

	fc := make(chan gox.FileChannel)
	defer close(fc)
	go fileService.Delete(fc, &mocks.Context{}, f)
	result := <- fc
	err := result.Err
	assert.Nil(err)
}

func TestRead(t *testing.T) {
	assert := assert.New(t)

	conn := &mocks.Databaser{}

	sn := "foo"
	pk := "0"
	f := helpers.ExampleFile()
	
	query := fmt.Sprintf("select * from read_file('%s')", f.FileId)
	var files []gox.File
	r := new(pg.Result)
	conn.On("Query", &files, query).Return(*r, nil)
	database := &postgres.Database{ Conn: conn }

	fileService := file.New(
		&mocks.Cacher{},
		database,
		&ffile.File{},
		&mocks.FileKinesis{},
		&faws.File{},
		sn,
		pk,
	)

	fc := make(chan gox.FileChannel)
	defer close(fc)

	go fileService.Read(fc, &mocks.Context{}, f)
	result := <- fc
	err := result.Err
	assert.Nil(err)
}

func TestReadAll(t *testing.T) {
	assert := assert.New(t)

	conn := &mocks.Databaser{}

	sn := "foo"
	pk := "0"
	
	query := "select * from read_all_files()"
	var files []gox.File
	r := new(pg.Result)
	conn.On("Query", &files, query).Return(*r, nil)
	database := &postgres.Database{ Conn: conn }

	fileService := file.New(
		&mocks.Cacher{},
		database,
		&ffile.File{},
		&mocks.FileKinesis{},
		&faws.File{},
		sn,
		pk,
	)

	fc := make(chan gox.FileChannel)
	defer close(fc)

	go fileService.ReadAll(fc, &mocks.Context{})
	result := <- fc
	err := result.Err
	assert.Nil(err)
}