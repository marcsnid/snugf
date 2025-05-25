package main

import (
	"io"
	"os"
	"testing"
)

func TestCompressFirst(t *testing.T) {
	transcodeKey := []byte(
		"hello friend this is a Great Key!@#%$#%#$^oeglo345623@#)()_+{}[]|;:/.,<>/?",
	)
	file, err := os.Open("README.md")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fileInfo, err := file.Stat()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	compressedReader, err := compressEncrypt(file, transcodeKey)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	finalReader, err := decryptDecompress(compressedReader, transcodeKey)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	finalB, err := io.ReadAll(finalReader)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if len(finalB) != int(fileInfo.Size()) {
		t.Fail()
		t.Log(
			"final bytes and file size not the same",
			len(finalB), fileInfo.Size())
	}
}
