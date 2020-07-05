package crypto

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestLinuxClientEncrypt(t *testing.T) {
	var encrypted string

	encrypted = LinuxClientEncrypt([]byte("hello"))
	assert.Equal(t, "7C9D21FDFE1C9186ECAB2A729F89A42D", encrypted, "Could not encrypt 'hello'.")

	encrypted = LinuxClientEncrypt([]byte("{ this is a long string ... }"))
	assert.Equal(t, "5D37AD47B8C6FF4A2CEB47A45BB021437A72886547FEA735105A409BEBC87F5B", encrypted, "Could not encrypt 'hello'.")

	encrypted = LinuxClientEncrypt([]byte("0123456789ABCDE"))
	assert.Equal(t, "8AF142630CD1590A8902E2CED105A625", encrypted, "Could not encrypt 15 char correctly.")
}

func TestEncryptAndroidRequestPayload(t *testing.T) {
	path := filepath.Join("testdata", t.Name()+".txt")
	content, err := ioutil.ReadFile(path)
	assert.Nil(t, err, "could not read test json: "+path)
	lines := strings.Split(strings.Replace(string(content), "\r", "", -1), "\n")

	url := lines[0]
	payload := lines[1]
	expected := lines[2]

	result := EncryptEAPIRequestPayload(url, []byte(payload))
	assert.Equal(t, expected, result, "incorrect payload generated.")
}
