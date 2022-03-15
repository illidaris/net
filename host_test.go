package net

import "testing"

func TestGetHostAddress(t *testing.T) {
	addr, err := GetHostAddress()
	if err != nil {
		t.Error(err)
	}
	t.Log("address is ", addr)
}

func TestGetHostAddresses(t *testing.T) {
	addrs, err := GetHostAddresses()
	if err != nil {
		t.Error(err)
	}
	for _, addr := range addrs {
		t.Log("address is ", addr)
	}
}

func BenchmarkGetHostAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addr, err := GetHostAddress()
		if err != nil {
			b.Error(err)
		}
		b.Log("address is ", addr)
	}
}
