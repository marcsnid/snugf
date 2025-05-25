package main

import (
	"compress/gzip"
	"io"

	"github.com/marcsnid/decouplet"
)

// Compresses the input io.Reader and returns a compressed io.Reader.
func compressBuffer(input io.Reader) (io.Reader, error) {
	pr, pw := io.Pipe()
	go func() {
		gw := gzip.NewWriter(pw)
		_, err := io.Copy(gw, input)
		gw.Close()
		pw.CloseWithError(err)
	}()
	return pr, nil
}

// compressEncrypt compresses the input io.Reader, encrypts it using the provided key,
// and then compresses it again. It returns a compressed io.Reader.
func compressEncrypt(input io.Reader, key []byte) (io.Reader, error) {
	// First compression stage
	compressR1, compressW1 := io.Pipe()
	go func() {
		gw := gzip.NewWriter(compressW1)
		_, err := io.Copy(gw, input)
		gw.Close()
		compressW1.CloseWithError(err)
	}()

	// Encryption stage
	encryptR, encryptW := io.Pipe()
	go func() {
		byteEncoder := decouplet.NewByteEncoder(key)
		err := byteEncoder.Encode(compressR1, encryptW)
		encryptW.CloseWithError(err)
	}()

	// Second compression stage
	compressR2, compressW2 := io.Pipe()
	go func() {
		gw := gzip.NewWriter(compressW2)
		_, err := io.Copy(gw, encryptR)
		gw.Close()
		compressW2.CloseWithError(err)
	}()

	return compressR2, nil
}

func decryptDecompress(input io.Reader, key []byte) (io.Reader, error) {
	// First decompression stage
	decompressR1, err := gzip.NewReader(input)
	if err != nil {
		return nil, err
	}

	// Decryption stage
	decryptR, decryptW := io.Pipe()
	go func() {
		byteEncoder := decouplet.NewByteEncoder(key)
		err := byteEncoder.Decode(decompressR1, decryptW)
		decryptW.CloseWithError(err)
	}()

	// Second decompression stage
	decompressR2, err := gzip.NewReader(decryptR)
	if err != nil {
		return nil, err
	}

	return decompressR2, nil
}
