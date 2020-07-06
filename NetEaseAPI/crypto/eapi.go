package crypto

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"
)

var eapiKey = []byte("e82ckenh8dichen8")

const eapiBoundary = "-36cd479b6b5-"

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
	data.WriteString(eapiBoundary)
	data.Write(input)
	data.WriteString(eapiBoundary)
	data.WriteString(hash)
	return data.Bytes()
}

func DecryptEAPIResponse(resp []byte) []byte {
	return decryptAESECB(resp, eapiKey)
}

func EncryptEAPIRequestPayload(path string, input []byte) string {
	hash := PrepareEAPIRequestHash(path, input)
	body := PrepareEAPIRequestBody(path, input, hash)

	return EncryptToString(body, eapiKey)
}

type EAPIRequestPayload struct {
	Path string
	Data string
}

func DecryptEAPIRequestBody(input []byte) (*EAPIRequestPayload, error) {
	decrypted := DecryptEAPIResponse(input)
	index := bytes.Index(decrypted, []byte(eapiBoundary))
	if index < 0 {
		return nil, errors.New("could not find request boundary")
	}
	return &EAPIRequestPayload{
		Path: string(decrypted[0:index]),
		Data: string(decrypted[index+len(eapiBoundary) : len(decrypted)-len(eapiBoundary)-32]),
	}, nil
}
