package main

import (
	"errors"
	"os"
)

func getFiles(filename string) ([]*os.File, error) {
	//TODO: will make this be able to use folders and have ability to do recursive reads
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, errors.New("cannot read folders")
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return []*os.File{file}, nil
}