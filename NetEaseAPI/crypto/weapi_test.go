package crypto

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWEAPISimpleRSA(t *testing.T) {
	var expected string

	expected = "86e36703b645357e675ae8b7816ec7271361688659cfb8bf813a8283b1ea1f4253cb85566e810e1dfddd18af5557bab967757" +
		"12636449e94d52888b49ff571a0c613f4f6da1192408f6499cb7ca5465a398cab18eb5702de37f5b12114304ca9505deea3d6971a939" +
		"775f74b8d52a4f8266f6a9c88771fae8eff46ea2d6bf133"
	crypted := weapiRSAEncrypt([]byte("hello"))
	assert.Equal(t, expected, strings.ToLower(hex.EncodeToString(crypted)))
}

func TestWEAPISimpleAES(t *testing.T) {
	var expected string

	expected = "tvMloIXulHxflDNlEEteTw=="
	crypted := encryptAESForWEAPI([]byte("hello"), []byte("1122334455667788"))
	assert.Equal(t, expected, crypted)
}

func TestEncryptWEAPIPayloadWithSecureKey(t *testing.T) {
	result, err := EncryptWEAPIPayloadWithSecureKey([]byte(`{"hello":1}`), []byte("bt6gZfVx7riGr82H"))
	assert.Nil(t, err)

	expectedSecureKey := "460ce3f5f2571867e7b734fa8dac4bdb9b934cc370eda6f9bb4df1a347ef414244c0d8805303f1809f95b716d2a" +
		"f40c5b356f07c013c054da9c46a1a4bbe0be3bf17c7c0812c41ae5b5acf0c843fe09f940bb4e143afd7bdbad0fb0b7313ece048f8f55" +
		"ad7db9cdfc98ef2cc5d4ba7eec976f7de2b9b7a40bf65d7be9d0f1f14"

	expectedParam := "nOZ7SPEtmdAx5BsuoM2DHXPZ2nfdYXZulEyogjkcsFE="

	assert.Equal(t, expectedSecureKey, result.SecureKey)
	assert.Equal(t, expectedParam, result.EncryptedText)
}
