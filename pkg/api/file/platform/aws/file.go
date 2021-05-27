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
	sn := aws.String(streamName)
	pk := aws.String(partitionKey)
	
	data, err := file.ToJson()
	if err != nil {
		return "", err
	}
	formatted, err := formatPayload(data, "create")
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

func (f File) Delete(kc FileKinesis, streamName string, fileId string, partitionKey string) (string, error){
	sn := aws.String(streamName)
	pk := aws.String(partitionKey)

	file := gox.File{FileId:fileId}
	data, err := file.ToJson()
	formatted, err := formatPayload(data, "delete")
	if err != nil {
		return "", err
	}
	output, err := kc.PutRecord(&kinesis.PutRecordInput {
		Data:			formatted,
		StreamName:		sn,
		PartitionKey:	pk,
	})
	if err != nil {
		return "", err
	}
	result := output.String()
	if err != nil {
		return "", err
	}

	return result, nil
}

func formatPayload(data string, operation string) ([]byte, error) {
	payload := make(map[string]string)
	payload["command"] = operation
	payload["data"] = data
	formatted, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return []byte(formatted), nil
}