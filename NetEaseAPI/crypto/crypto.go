package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"strings"
)

func PKCS7Padding(data []byte, blockSize int) []byte {
	paddingSize := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, padding...)
}

func PKCS7UnPadding(origData []byte) []byte {
	dataSize := len(origData)
	paddingSize := int(origData[dataSize-1])
	return origData[:(dataSize - paddingSize)]
}

func encryptAESECB(plaintext, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	plaintext = PKCS7Padding(plaintext, block.BlockSize())
	encrypted := make([]byte, len(plaintext))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(plaintext); bs, be = bs+size, be+size {
		block.Encrypt(encrypted[bs:be], plaintext[bs:be])
	}

	return encrypted
}

func encryptAESCBC(plaintext, key, iv []byte) []byte {
	block, _ := aes.NewCipher(key)
	plaintext = PKCS7Padding(plaintext, block.BlockSize())
	encrypted := make([]byte, len(plaintext))

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(encrypted, plaintext)
	return encrypted
}

func decryptAESECB(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EncryptToString(input, key []byte) string {
	encrypted := encryptAESECB(input, key)
	return strings.ToUpper(hex.EncodeToString(encrypted))
}
