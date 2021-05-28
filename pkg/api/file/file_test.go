package file_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/stretchr/testify/assert"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/mocks"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	"github.com/konstantinfarrell/go-example/pkg/api/file"
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
	database := &postgres.Database{ Conn: &mocks.Databaser{}}

	sn := "foo"
	pk := "0"
	f := exampleFile()
	data, _ := f.ToJson()
	formatted, _ := faws.FormatPayload(data, "create")
	payload := &kinesis.PutRecordInput {
		Data:			formatted,
		StreamName:		&sn,
		PartitionKey:	&pk,
	}
	ks := &mocks.FileKinesis{}
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
	go fileService.Create(fc, &mocks.Context{}, f)
	result := <- fc
	err := result.Err

	assert.Equal(t, err, nil)
}