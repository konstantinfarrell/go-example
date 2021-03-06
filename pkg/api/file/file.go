package file

import (
	"github.com/konstantinfarrell/go-example"
)

// Creates and deletes are sent to kinesis
func (f File) Create(fc chan gox.FileChannel, c Context, request *gox.File) {
	_, err := f.aws.Create(f.kinesis, f.streamName, request, f.partitionKey)
	if err != nil {
		fc <- gox.FileChannel{File: nil, Err: err}
	}
	result := []gox.File{*request}
	fc <- gox.FileChannel{File: &result, Err: nil}
}

func (f File) Delete(fc chan gox.FileChannel, c Context, request *gox.File) {
	_, err := f.aws.Delete(f.kinesis, f.streamName, request.FileId, f.partitionKey)
	if err != nil {
		fc <- gox.FileChannel{File: nil, Err: err}
	}
	fc <- gox.FileChannel{File: nil, Err: nil}
}

// Reads are sent to the DB
func (f File) Read(fc chan gox.FileChannel, c Context, request *gox.File) {
	result, err := f.database.ReadFile(f.pgs, request)
	if err != nil {
		fc <- gox.FileChannel{File: nil, Err: err}
	}

	if result == nil {
		fc <- gox.FileChannel{File: nil, Err: err}
		return
	}

	fc <- gox.FileChannel{File: result, Err: nil}
}

func (f File) ReadAll(fc chan gox.FileChannel, c Context) {
	result, err := f.database.ReadAllFiles(f.pgs)
	if err != nil {
		fc <- gox.FileChannel{File: nil, Err: err}
	}
	
	fc <- gox.FileChannel{File: result, Err: nil}
}

func (f File) GetResult(fc chan gox.FileChannel) *gox.FileChannel {
	// TODO: idk maybe timeout or something? still learning
	result := <- fc
	return &result
}