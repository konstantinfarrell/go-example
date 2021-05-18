package file

import (
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/labstack/echo"
	rs "github.com/go-redis/redis"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
	faws "github.com/konstantinfarrell/go-example/pkg/api/file/platform/aws"
	ffile "github.com/konstantinfarrell/go-example/pkg/api/file/platform/postgres"
)

type Service interface {
	Create(echo.Context, *gox.File) (*gox.File, error)
	ReadAll(echo.Context) (*[]gox.File, error)
	Read(echo.Context, *gox.File) (*gox.File, error)
	Delete(echo.Context, *gox.File) error
}

func New(cache *rs.Client, pgs *postgres.Database, database *ffile.File, kinesis *kinesis.Kinesis, aws *faws.File, streamName string, partitionKey string) *File {
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

func Initialize(cache *rs.Client, pgs *postgres.Database, database *ffile.File, kinesis *kinesis.Kinesis, aws *faws.File, streamName string, partitionKey string) *File {
	return New(cache, pgs, database, kinesis, aws, streamName, partitionKey)
}

// File represents file application service
type File struct {
	cache 			*rs.Client
	pgs				*postgres.Database
	database		*ffile.File
	aws				*faws.File
	kinesis			*kinesis.Kinesis
	streamName		string
	partitionKey	string
}