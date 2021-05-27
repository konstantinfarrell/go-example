package file

import (
	"github.com/labstack/echo"
	rs "github.com/go-redis/redis"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

type Service interface {
	Create(chan gox.FileChannel, echo.Context, *gox.File)
	ReadAll(chan gox.FileChannel, echo.Context)
	Read(chan gox.FileChannel, echo.Context, *gox.File) 
	Delete(chan gox.FileChannel, echo.Context, *gox.File)
}

func New(cache *rs.Client, pgs *postgres.Database, database *ffile.File, kinesis faws.FileKinesis, aws *faws.File, streamName string, partitionKey string) *File {
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

func Initialize(cache *rs.Client, pgs *postgres.Database, database *ffile.File, kinesis faws.FileKinesis, aws *faws.File, streamName string, partitionKey string) *File {
	return New(cache, pgs, database, kinesis, aws, streamName, partitionKey)
}




// File represents file application service
type File struct {
	cache 			*rs.Client
	pgs				*postgres.Database
	database		*ffile.File
	aws				*faws.File
	kinesis			faws.FileKinesis
	streamName		string
	partitionKey	string
}