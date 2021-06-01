package aws_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/konstantinfarrell/go-example/mocks"
	"github.com/konstantinfarrell/go-example/pkg/util/helpers"
	"github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
)

func TestCreate(t *testing.T){
	assert := assert.New(t)

	faws := new(aws.File)
	f := helpers.ExampleFile()
	kc := new(mocks.FileKinesis)
	kc.On("PutRecord", mock.Anything).Return(new(kinesis.PutRecordOutput), nil)

	_, err := faws.Create(kc, "", f, "")
	assert.Nil(err)	
}

func TestDelete(t *testing.T){
	assert := assert.New(t)

	faws := new(aws.File)
	f := helpers.ExampleFile()
	kc := new(mocks.FileKinesis)
	kc.On("PutRecord", mock.Anything).Return(new(kinesis.PutRecordOutput), nil)

	_, err := faws.Delete(kc, "", f.FileId, "")
	assert.Nil(err)	
}