package file

import (
	"github.com/labstack/echo"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	"github.com/konstantinfarrell/go-example/pkg/util/redis"
	faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

type Context interface {
	echo.Context
}

type Service interface {
	Create(chan gox.FileChannel, Context, *gox.File)
	ReadAll(chan gox.FileChannel, Context)
	Read(chan gox.FileChannel, Context, *gox.File) 
	Delete(chan gox.FileChannel, Context, *gox.File)
	GetResult(chan gox.FileChannel) *gox.FileChannel
}

func New(cache redis.Cacher, pgs *postgres.Database, database *ffile.File, kinesis faws.FileKinesis, aws *faws.File, streamName string, partitionKey string) *File {
	return &File{
		cache: cache,
		pgs: pgs,
		database: database,
		kinesis: kinesis,
		aws: aws,
		streamName: streamName,
		partitionKey: partitionKey,
	}
}

func Initialize(cache redis.Cacher, pgs *postgres.Database, database *ffile.File, kinesis faws.FileKinesis, aws *faws.File, streamName string, partitionKey string) *File {
	return New(cache, pgs, database, kinesis, aws, streamName, partitionKey)
}


// File represents file application service
type File struct {
	cache 			redis.Cacher
	pgs				*postgres.Database
	database		*ffile.File
	aws				*faws.File
	kinesis			faws.FileKinesis
	streamName		string
	partitionKey	string
}