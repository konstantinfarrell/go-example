package config

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"

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

	var s Server
	err = envconfig.Process("", &s)
	if err != nil {
		return nil, fmt.Errorf("Error reading environment variables, %v", err.Error())
	}
	conf.Server = &s

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
	Port 			string 	`envconfig:"API_PORT"`
	Loglevel 		string	`envconfig:"API_LOG_LEVEL"`
	TimeoutSeconds 	int		`envconfig:"API_TIMEOUT`
}

type Aws struct {
	AWSAccessKeyId		string	`envconfig:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey 	string	`envconfig:"AWS_SECRET_ACCESS_KEY"`	
	Region				string	`envconfig:"AWS_REGION"`
	KinesisStreamName	string	`envconfig:"KINESIS_STREAM_NAME"`
	PartitionKey		string	`envconfig:"KINESIS_PARTITION_KEY"`
}