package ip

import "testing"

func TestIpFormat(t *testing.T) {
	var ip string

	ip = Format(Make(255, 255, 255, 255))
	if ip != "255.255.255.255" {
		t.Errorf("could not make 255.255.255.255, got %s instead.", ip)
	}

	ip = Format(Make(1, 3, 5, 7))
	if ip != "1.3.5.7" {
		t.Errorf("could not make 1.3.5.7, got %s instead.", ip)
	}
}
