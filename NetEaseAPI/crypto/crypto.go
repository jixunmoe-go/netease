package crypto

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"strings"
)

var secret = []byte("rFgB&h#%2?^eDg:Q")

func PKCS7Padding(data []byte, blockSize int) []byte {
	paddingSize := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(data, padding...)
}

func encryptEcb(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func Encrypt(data []byte) []byte {
	return encryptEcb(data, secret)
}

func EncryptString(str string) string {
	encrypted := Encrypt([]byte(str))
	return strings.ToUpper(hex.EncodeToString(encrypted))
}

func EncryptToString(str []byte) string {
	encrypted := Encrypt(str)
	return strings.ToUpper(hex.EncodeToString(encrypted))
}
