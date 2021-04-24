package parse

import (
	"fmt"
	"net"
)

func ParseIp(value string) (net.IP, error) {
	ip := net.ParseIP(value)

	if ip == nil {
		return nil, fmt.Errorf("parseIp: '%s' is no valid IP", value)
	}

	return ip, nil
}
