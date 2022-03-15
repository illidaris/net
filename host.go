package net

import (
	"errors"
	"fmt"
	"net"
)

// GetHostAddressFromDNS selects a valid outbound IP address for the host.
func GetHostAddressFromDNS(dns string) (string, error) {
	// Use udp so no handshake is made.
	// Any IP can be used, since connection is not established, but we used a known DNS IP.
	conn, err := net.Dial("udp", fmt.Sprintf("%s:80", dns))
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String(), nil
}

// GetHostAddress selects a valid outbound IP address for the host.
func GetHostAddress() (string, error) {
	// Could not find one via a  UDP connection, so we fallback to the "old" way: try first non-loopback IPv4:
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("could not determine host IP address")
}

// GetHostAddresses select all ip address
func GetHostAddresses() ([]string, error) {
	addresses := make([]string, 0)
	// Could not find one via a  UDP connection, so we fallback to the "old" way: try first non-loopback IPv4:
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return addresses, err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addresses = append(addresses, ipnet.IP.String())
			}
		}
	}
	return addresses, nil
}
