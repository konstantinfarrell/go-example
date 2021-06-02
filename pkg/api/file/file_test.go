package file_test

import (
	"fmt"
	"testing"

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
		new(mocks.Cacher),
		database,
		new(ffile.File),
		ks,
		new(faws.File),
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
		new(mocks.Cacher),
		database,
		new(ffile.File),
		ks,
		new(faws.File),
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
		new(mocks.Cacher),
		database,
		new(ffile.File),
		new(mocks.FileKinesis),
		new(faws.File),
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

	conn := new(mocks.Databaser)

	sn := "foo"
	pk := "0"
	
	query := "select * from read_all_files()"
	var files []gox.File
	r := new(pg.Result)
	conn.On("Query", &files, query).Return(*r, nil)
	database := &postgres.Database{ Conn: conn }

	fileService := file.New(
		new(mocks.Cacher),
		database,
		new(ffile.File),
		new(mocks.FileKinesis),
		new(faws.File),
		sn,
		pk,
	)

	fc := make(chan gox.FileChannel)
	defer close(fc)

	go fileService.ReadAll(fc, new(mocks.Context))
	result := <- fc
	err := result.Err
	assert.Nil(err)
}