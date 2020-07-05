package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var linuxAPIKey = []byte("rFgB&h#%2?^eDg:Q")
var eapiKey = []byte("e82ckenh8dichen8")

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

func decryptAESECB(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func Encrypt(input, key []byte) string {
	encrypted := encryptAESECB(input, key)
	return strings.ToUpper(hex.EncodeToString(encrypted))
}

func LinuxClientEncrypt(input []byte) string {
	return Encrypt(input, linuxAPIKey)
}

func PrepareEAPIRequestHash(url string, input []byte) string {
	hash := md5.New()
	hash.Write([]byte("nobody"))
	hash.Write([]byte(url))
	hash.Write([]byte("use"))
	hash.Write(input)
	hash.Write([]byte("md5forencrypt"))
	return strings.ToLower(hex.EncodeToString(hash.Sum(nil)))
}

func PrepareEAPIRequestBody(url string, input []byte, hash string) []byte {
	data := bytes.Buffer{}
	data.WriteString(url)
	data.WriteString("-36cd479b6b5-")
	data.Write(input)
	data.WriteString("-36cd479b6b5-")
	data.WriteString(hash)
	return data.Bytes()
}

func DecryptEAPIResponse(resp []byte) []byte {
	return decryptAESECB(resp, eapiKey)
}

func EncryptEAPIRequestPayload(path string, input []byte) string {
	hash := PrepareEAPIRequestHash(path, input)
	body := PrepareEAPIRequestBody(path, input, hash)

	return Encrypt(body, eapiKey)
}
