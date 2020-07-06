package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"math/big"
	"math/rand"
)

var weapiAESKey = []byte("0CoJUm6Qyw8W8jud")
var weapiAESKeyIV = []byte("0102030405060708")

var weapiRSAEncryptionExpBigInt *big.Int
var weapiRSAModulusBigInt *big.Int
var weapiRSAModulus = []byte{
	0xE0, 0xB5, 0x09, 0xF6, 0x25, 0x9D, 0xF8, 0x64,
	0x2D, 0xBC, 0x35, 0x66, 0x29, 0x01, 0x47, 0x7D,
	0xF2, 0x26, 0x77, 0xEC, 0x15, 0x2B, 0x5F, 0xF6,
	0x8A, 0xCE, 0x61, 0x5B, 0xB7, 0xB7, 0x25, 0x15,
	0x2B, 0x3A, 0xB1, 0x7A, 0x87, 0x6A, 0xEA, 0x8A,
	0x5A, 0xA7, 0x6D, 0x2E, 0x41, 0x76, 0x29, 0xEC,
	0x4E, 0xE3, 0x41, 0xF5, 0x61, 0x35, 0xFC, 0xCF,
	0x69, 0x52, 0x80, 0x10, 0x4E, 0x03, 0x12, 0xEC,
	0xBD, 0xA9, 0x25, 0x57, 0xC9, 0x38, 0x70, 0x11,
	0x4A, 0xF6, 0xC9, 0xD0, 0x5C, 0x4F, 0x7F, 0x0C,
	0x36, 0x85, 0xB7, 0xA4, 0x6B, 0xEE, 0x25, 0x59,
	0x32, 0x57, 0x5C, 0xCE, 0x10, 0xB4, 0x24, 0xD8,
	0x13, 0xCF, 0xE4, 0x87, 0x5D, 0x3E, 0x82, 0x04,
	0x7B, 0x97, 0xDD, 0xEF, 0x52, 0x74, 0x1D, 0x54,
	0x6B, 0x8E, 0x28, 0x9D, 0xC6, 0x93, 0x5B, 0x3E,
	0xCE, 0x04, 0x62, 0xDB, 0x0A, 0x22, 0xB8, 0xE7,
}

func setupKeyForWEAPI() {
	if weapiRSAModulusBigInt == nil {
		weapiRSAModulusBigInt = &big.Int{}
		weapiRSAModulusBigInt.SetBytes(weapiRSAModulus)
		weapiRSAEncryptionExpBigInt = big.NewInt(0x010001)
	}
}

func weapiRSAEncrypt(input []byte) []byte {
	setupKeyForWEAPI()

	m := big.Int{}
	c := big.Int{}

	size := len(input)
	reversed := make([]byte, size)
	for i, j := 0, size-1; i <= j; i, j = i+1, j-1 {
		reversed[i] = input[j]
		reversed[j] = input[i]
	}
	m.SetBytes(reversed)

	c.Exp(&m, weapiRSAEncryptionExpBigInt, weapiRSAModulusBigInt)
	crypted := c.Bytes()

	if len(crypted) < 128 {
		crypted = append(bytes.Repeat([]byte{0}, 128-len(crypted)), crypted...)
	}

	return crypted
}

type WEAPICryptedPayload struct {
	EncryptedText string `json:"params"`
	SecureKey     string `json:"encSecKey"`
}

func encryptAESForWEAPI(input, key []byte) string {
	crypted := encryptAESCBC(input, key, weapiAESKeyIV)
	return base64.StdEncoding.EncodeToString(crypted)
}

func EncryptWEAPIPayloadWithSecureKey(input, secureKey []byte) (*WEAPICryptedPayload, error) {
	crypted := encryptAESForWEAPI(input, weapiAESKey)
	encText := encryptAESForWEAPI([]byte(crypted), secureKey)
	encSecKey := weapiRSAEncrypt(secureKey)

	var result = &WEAPICryptedPayload{}
	result.EncryptedText = encText
	result.SecureKey = hex.EncodeToString(encSecKey)
	return result, nil
}

const charMap = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateSecureKey() []byte {
	sb := bytes.Buffer{}
	sb.Grow(16)
	for i := 0; i < 16; i++ {
		sb.WriteByte(charMap[rand.Intn(62)])
	}
	return sb.Bytes()
}

func EncryptWEAPIPayload(input []byte) (*WEAPICryptedPayload, error) {
	var secureKey = generateSecureKey()
	return EncryptWEAPIPayloadWithSecureKey(input, secureKey)
}
