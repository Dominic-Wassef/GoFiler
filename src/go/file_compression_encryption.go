package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

const (
	key = "myverystrongpasswordo32bitlength"
)

// Compress compresses the input bytes using gzip.
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress data: %w", err)
	}

	if err := gz.Close(); err != nil {
		return nil, fmt.Errorf("failed to close gzip writer: %w", err)
	}

	return buf.Bytes(), nil
}

// Decompress decompresses the input bytes using gzip.
func Decompress(data []byte) ([]byte, error) {
	rdata := bytes.NewReader(data)
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}

	var res bytes.Buffer
	_, err = res.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return res.Bytes(), nil
}

// Encrypt encrypts the input data using AES.
func Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to create nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt decrypts the input data using AES.
func Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}

// CompressAndEncrypt compresses and then encrypts the input data.
func CompressAndEncrypt(data []byte) ([]byte, error) {
	compressedData, err := Compress(data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress data: %w", err)
	}

	encryptedData, err := Encrypt(compressedData)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data: %w", err)
	}

	return encryptedData, nil
}

// DecryptAndDecompress decrypts and then decompresses the input data.
func DecryptAndDecompress(data []byte) ([]byte, error) {
	decryptedData, err := Decrypt(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data: %w", err)
	}

	decompressedData, err := Decompress(decryptedData)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %w", err)
	}

	return decompressedData, nil
}

// CompressAndEncryptFile reads the file, compresses and encrypts its content, and writes it back to the file.
func CompressAndEncryptFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	encryptedData, err := CompressAndEncrypt(data)
	if err != nil {
		return fmt.Errorf("failed to compress and encrypt file: %w", err)
	}

	err = os.WriteFile(filename, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// DecryptAndDecompressFile reads the file, decrypts and decompresses its content, and writes it back to the file.
func DecryptAndDecompressFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	decryptedData, err := DecryptAndDecompress(data)
	if err != nil {
		return fmt.Errorf("failed to decrypt and decompress file: %w", err)
	}

	err = os.WriteFile(filename, decryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}