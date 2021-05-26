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
	// this was nice when config files were the way to go
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
	var aws Aws
	err := envconfig.Process("", &aws)
	if err != nil {
		return nil, fmt.Errorf("Error reading environment variables, %v", err.Error())
	}
	conf.AWS = &aws

	var db Database 
	err = envconfig.Process("", &db)
	if err != nil {
		return nil, fmt.Errorf("Error reading environment variables, %v", err.Error())
	}
	conf.DB = &db

	var c Cache
	err = envconfig.Process("", &c)
	if err != nil {
		return nil, fmt.Errorf("Error reading environment variables, %v", err.Error())
	}
	conf.Cache = &c

	return conf, nil
}

type Configuration struct {
	DB *Database
	Server *Server
	AWS *Aws
	Cache *Cache
}

type Database struct {
	User string		`envconfig:"DB_USER"`
	Pass string		`envconfig:"DB_PASS"`
	Name string		`envconfig:"DB_NAME"`	
	Port int		`envconfig:"DB_PORT"`
	Addr string		`envconfig:"DB_ADDR"`
}

type Cache struct {
	Addr 	string	`envconfig:"CACHE_ADDR"`
	Pass 	string	`envconfig:"CACHE_PASS"`
	DB		int		`envconfig:"CACHE_DB"`
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
	Region				string	`envconfig:"AWS_REGION"`
	KinesisStreamName	string	`envconfig:"KINESIS_STREAM_NAME"`
	PartitionKey		string	`envconfig:"KINESIS_PARTITION_KEY"`
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