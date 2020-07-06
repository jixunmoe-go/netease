package crypto

var linuxAPIKey = []byte("rFgB&h#%2?^eDg:Q")

func LinuxClientEncrypt(input []byte) string {
	return Encrypt(input, linuxAPIKey)
}
