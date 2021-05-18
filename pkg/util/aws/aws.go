package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func New(region string) (*kinesis.Kinesis, error){
	s := session.New(&aws.Config{
		Region: aws.String(region),
	})
	kc := kinesis.New(s)
	return kc, nil
}