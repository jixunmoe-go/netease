package crypto

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	var encrypted string

	encrypted = EncryptString("hello")
	if encrypted != "7C9D21FDFE1C9186ECAB2A729F89A42D" {
		t.Error("Could not encrypt 'hello'.")
	}

	encrypted = EncryptString("{ this is a long string ... }")
	if encrypted != "5D37AD47B8C6FF4A2CEB47A45BB021437A72886547FEA735105A409BEBC87F5B" {
		t.Error("Could not encrypt 'hello'.")
	}

	encrypted = EncryptString("0123456789ABCDE")
	if encrypted != "8AF142630CD1590A8902E2CED105A625" {
		t.Error("Could not encrypt 15 char correctly.")
	}
}
