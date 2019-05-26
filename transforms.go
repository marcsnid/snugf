package main

import (
	"compress/gzip"
	"github.com/marcsj/decouplet"
	"io"
)

func compressBuffer(input io.Reader) (*io.PipeReader, error) {
	reader, writer := io.Pipe()
	go func() {
		defer writer.Close()

		zipWriter, err := gzip.NewWriterLevel(writer, gzip.BestCompression)
		defer zipWriter.Close()
		if err != nil {
			writer.CloseWithError(err)
		}

		_, err = io.Copy(zipWriter, input)
		if err != nil {
			writer.CloseWithError(err)
		}
	}()
	return reader, nil
}

func compressEncrypt(input io.Reader, key []byte) (*io.PipeReader, error) {
	compressR, err := compressBuffer(input)
	if err != nil {
		return nil, err
	}
	decoupledReader := decouplet.EncodeBytesStream(
		compressR, key)
	if err != nil {
		return nil, err
	}
	secondCompressR, err := compressBuffer(decoupledReader)
	if err != nil {
		return nil, err
	}
	return secondCompressR, err
}

func decryptDecompress(input io.Reader, key []byte) (io.Reader, error) {
	decompressedR, err := gzip.NewReader(input)
	if err != nil {
		return nil, err
	}
	decodedR, err := decouplet.DecodeBytesStream(decompressedR, key)
	if err != nil {
		return nil, err
	}
	secondDecompressedR, err := gzip.NewReader(decodedR)
	if err != nil {
		return nil, err
	}
	return secondDecompressedR, nil
}