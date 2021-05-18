package postgres

import (
	"github.com/konstantinfarrell/go-example"
	"github.com/konstantinfarrell/go-example/pkg/util/postgres"
)

type File struct {}

func (f *File) ReadFile(d *postgres.Database, file *gox.File) (*gox.File, error) {
	spname := "read_file"
	var files []gox.File
	_, err := d.Call(true, &files, spname, file.FileId)
	if err != nil {
		return nil, err
	}
	return &files[0], nil
}

func (f *File) ReadAllFiles(d *postgres.Database) (*[]gox.File, error) {
	spname := "read_all_files"
	var files []gox.File
	_, err := d.Call(true, &files, spname)
	if err != nil {
		return nil, err
	}
	return &files, nil
}