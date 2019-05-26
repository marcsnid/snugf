package main

import (
	"errors"
	"github.com/urfave/cli"
	"io/ioutil"
)

type operator struct {
	fileKey *string
	stringKey *string
}

func (o operator) getKey() ([]byte, error){
	if *o.fileKey != "" {
		file, err := ioutil.ReadFile(*o.fileKey)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else if *o.stringKey != "" {
		return []byte(*o.stringKey), nil
	}
	return nil, errors.New("no key in use")
}

func (o operator) write(c *cli.Context) error {
	fileRead := c.Args().Get(0)
	fileWritten := c.Args().Get(1)
	key, err := o.getKey()
	if err != nil {
		return err
	}

	files, err := getFiles(fileRead)
	if err != nil {
		return err
	}
	compressedReader, err := compressEncrypt(files[0], key)
	if err != nil {
		return err
	}
	defer compressedReader.Close()

	return writeFile(fileWritten, compressedReader)
}

func (o operator) read(c *cli.Context) error {
	fileRead := c.Args().Get(0)
	fileWritten := c.Args().Get(1)
	key, err := o.getKey()
	if err != nil {
		return err
	}

	files, err := getFiles(fileRead)
	if err != nil {
		return err
	}
	compressedReader, err := decryptDecompress(files[0], key)
	if err != nil {
		return err
	}
	return writeFile(fileWritten, compressedReader)
}