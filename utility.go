package main

import (
	"bufio"
	"errors"
	"io"
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

func writeFile(filename string, input io.Reader) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = io.Copy(w, input)
	if err != nil {
		return err
	}

	return w.Flush()
}
