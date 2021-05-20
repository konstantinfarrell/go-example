package gox // LOL mtgox

import (
	"encoding/json"
)

type FileChannel struct {
	File 	*[]File
	Err		error
} 

type File struct {
	FileId		string
	Filename 	string
	Path		string
	Permissions string
	Created		string
	Modified	string
	Data		[]byte
	Received	string
}

func (f *File) ToJson() (string, error) {
	file, err := json.Marshal(f)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

func (f *File) FromJson(data string) (*File, error) {
	err := json.Unmarshal([]byte(data), f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *File) FromJsonMany(data string) ([]File, error) {
	var files []File
	err := json.Unmarshal([]byte(data), files)
	if err != nil {
		return nil, err
	}
	return files, nil
}