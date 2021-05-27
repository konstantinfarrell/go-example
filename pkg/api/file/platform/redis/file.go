package redis

import (
	"context"
	"fmt"

	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/redis"
)

var FileKey = "f:%s"
var ctx = context.Background()

type File struct{}

func (f File) Create(client redis.Cacher, file gox.File) (*gox.File, error){
	var key = fmt.Sprintf(FileKey, file.FileId)
	jsonFile, err := file.ToJson()
	if err != nil {
		return nil, err
	}

	err = client.Set(ctx, key, jsonFile, 0).Err()

	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (f File) Read(client redis.Cacher, fileId string) (*gox.File, error) {
	var key = fmt.Sprintf(FileKey, fileId)
	var file = new(gox.File)
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	jsonFile, err := file.FromJson(val)
	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}

func (f File) ReadAll(client redis.Cacher) (*[]gox.File, error) {
	var key = fmt.Sprintf(FileKey, "*")
	file := new(gox.File)
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	jsonFiles, err := file.FromJsonMany(val)
	if err != nil {
		return nil, err
	}

	return &jsonFiles, nil
}

func (f File) Delete(client redis.Cacher, file gox.File) error {
	var key = fmt.Sprintf(FileKey, file.FileId)
	err := client.Del(ctx, key).Err()
	return err
}