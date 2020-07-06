package crypto

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

var eapiKey = []byte("e82ckenh8dichen8")

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
