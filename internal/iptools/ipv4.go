package iptools

import (
	"net"
)

func GetIpv4sByDomain(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}

	var ret []net.IP
	for _, ip := range ips {
		ipV4 := ip.To4()
		if ipV4 != nil {
			ret = append(ret, ipV4)
		}
	}

	return ret, nil
}

func IsValidIPv4(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}

	if ip.To4() == nil {
		return false
	}

	return true
}
