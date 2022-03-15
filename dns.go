package net

type DNS string

var (
	GoogleDNS  DNS = "8.8.8.8"
	TelecomDNS DNS = "114.114.114.114"
	BaiduDNS   DNS = "180.76.76.76"
	AlibabaDNS DNS = "223.5.5.5"
)

func (d DNS) GetHostAddress() (string, error) {
	return GetHostAddressFromDNS(string(d))
}
