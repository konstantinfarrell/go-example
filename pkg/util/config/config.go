package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/konstantinfarrell/go-example/pkg/util/helpers"
)

func Load(path string) (*Configuration, error) {
	var conf = new(Configuration)
	conf, err := LoadFromConfig(path, conf)
	if err != nil {
		return nil, err
	}

	conf, err = LoadFromEnvVar(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func LoadFromConfig(path string, conf *Configuration) (*Configuration, error){
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file, %s", err)
	}
	if err:= yaml.Unmarshal(bytes, conf); err != nil {
		return nil, fmt.Errorf("Unable to decode into struct, %v", err)
	}
	return conf, nil
}

func LoadFromEnvVar(conf *Configuration) (*Configuration, error) {
	var aws Aws // Not necessary since AWS libs read from env vars anyway but still a good exercise
	err := envconfig.Process("", &aws)
	if err != nil {
		return nil, fmt.Errorf("Error reading environment variables, %v", err.Error())
	}
	conf.AWS.AWSAccessKeyId = aws.AWSAccessKeyId
	conf.AWS.AWSSecretAccessKey = aws.AWSSecretAccessKey
	return conf, nil
}

type Configuration struct {
	DB *Database
	Server *Server
	AWS *Aws
	Cache *Cache
}

type Database struct {
	User string
	Pass string
	Name string
	Port int
	Addr string
}

type Cache struct {
	Addr 	string
	Pass 	string
	DB		int
}

type Server struct {
	Port 			string
	Loglevel 		string
	TimeoutSeconds 	int		`yaml:"timeout_seconds"`
}

// Unnecessary since aws packages read from environemnt variables
type Aws struct {
	AWSAccessKeyId		string	`envconfig:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey 	string	`envconfig:"AWS_SECRET_ACCESS_KEY"`	
	Region				string	
	KinesisStreamName	string	`yaml:"kinesis_stream_name"`
	PartitionKey		string	`yaml:"partition_key"`
}

func printTypes(item Configuration){
	itemVal := reflect.ValueOf(item)
    for i := 0; i < itemVal.NumField(); i++ {
        fieldVal := itemVal.Field(i)
        if fieldVal.Kind() == reflect.Ptr {
            fieldVal = fieldVal.Elem()
        }
        fmt.Println(fieldVal.Kind())
    }
}