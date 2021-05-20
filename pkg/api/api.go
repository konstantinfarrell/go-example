package api

import (
	_ "crypto/sha1"
	_ "os"

	"github.com/konstantinfarrell/go-example/pkg/util/config"
	"github.com/konstantinfarrell/go-example/pkg/util/server"
	"github.com/konstantinfarrell/go-example/pkg/util/redis"
	"github.com/konstantinfarrell/go-example/pkg/util/aws"
	"github.com/konstantinfarrell/go-example/pkg/util/zlog"
	ps "github.com/konstantinfarrell/go-example/pkg/util/postgres"

	"github.com/konstantinfarrell/go-example/pkg/api/file"
	ft "github.com/konstantinfarrell/go-example/pkg/api/file/transport"
	fl "github.com/konstantinfarrell/go-example/pkg/api/file/logging"
	
	faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

func Start(conf *config.Configuration) error {
	// Initialize services from configuration object

	// Cache
	cache, err := redis.New(conf.Cache.Addr, conf.Cache.Pass, conf.Cache.DB)
	if err != nil {
		panic(err)	
	}
	
	// Database
	pg, err := ps.New(conf)
	if err != nil {
		panic(err)
	}
	database := new(ffile.File)

	// AWS kinesis
	kinesis, err := aws.New(conf.AWS.Region)
	if err != nil {
		panic(err)
	}

	streamName := conf.AWS.KinesisStreamName
	partitionKey := conf.AWS.PartitionKey
	aws := new(faws.File)

	// Logging
	log := zlog.New()

	// Instantiate server
	e := server.New()
	
	// Setup base path
	v1 := e.Group("/v1")

	// Inject initialized services into API
	ft.NewHTTP(fl.New(file.Initialize(cache, pg, database, kinesis, aws, streamName, partitionKey), log), v1)

	configuration := &server.Config{
		Port:			conf.Server.Port,
		LogLevel:		conf.Server.Loglevel,
		TimeoutSeconds: conf.Server.TimeoutSeconds,
	}

	server.Start(e, configuration)

	return nil
}