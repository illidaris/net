package net

import (
	"fmt"
	"testing"
	"time"
)

func TestDNS(t *testing.T) {
	for _, dns := range []DNS{GoogleDNS, TelecomDNS, BaiduDNS, AlibabaDNS} {
		begin := time.Now()
		addr, err := dns.GetHostAddress()
		if err != nil {
			t.Error(err)
		}
		t.Log(fmt.Sprintf("[%s]address is %s cost %dns", dns, addr, time.Since(begin).Nanoseconds()))
	}
}
