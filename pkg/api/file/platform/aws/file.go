package aws

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"

	"github.com/konstantinfarrell/go-example"
)


type FileKinesis interface {
	PutRecord(*kinesis.PutRecordInput) (*kinesis.PutRecordOutput, error)
}

type File struct{}

func (f File) Create(kc FileKinesis, streamName string, file *gox.File, partitionKey string) (string, error){
	return f.sendPayload(kc, streamName, file, partitionKey, "create")
}

func (f File) Delete(kc FileKinesis, streamName string, fileId string, partitionKey string) (string, error){
	file := &gox.File{FileId:fileId}
	return f.sendPayload(kc, streamName, file, partitionKey, "delete")
}

func (f File) sendPayload(kc FileKinesis, streamName string, file *gox.File, partitionKey string, operation string) (string, error) {
	sn := aws.String(streamName)
	pk := aws.String(partitionKey)
	
	data, err := file.ToJson()
	if err != nil {
		return "", err
	}
	formatted, err := FormatPayload(data, operation)
	if err != nil {
		return "", err
	}

	output, err := kc.PutRecord(&kinesis.PutRecordInput {
		Data:			formatted,
		StreamName:		sn,
		PartitionKey:	pk,
	})
	result := output.String()
	if err != nil {
		return "", err
	}
	return result, nil
}

func FormatPayload(data string, operation string) ([]byte, error) {
	payload := make(map[string]string)
	payload["command"] = operation
	payload["data"] = data
	formatted, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return []byte(formatted), nil
}