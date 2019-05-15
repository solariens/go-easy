package utils

import (
	"net"
	"sync"
)

var (
	mapServerIP  = make(map[string]string)
	serverIPLock = sync.RWMutex{}
)

func ParseServerIP(eth string) string {
	serverIPLock.RLock()
	if ip, ok := mapServerIP[eth]; ok {
		serverIPLock.RUnlock()
		return ip
	}

	serverIPLock.RUnlock()

	ip := GetEthIP(eth)
	serverIPLock.Lock()
	mapServerIP[eth] = ip
	serverIPLock.Unlock()

	return ip
}

func GetEthIP(eth string) string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return eth
	}

	for _, item := range ifaces {
		addrs, err := item.Addrs()
		if err == nil && item.Name == eth {
			for _, addr := range addrs {
				var ip net.IP
				switch addr.(type) {
				case *net.IPNet:
					ip = addr.(*net.IPNet).IP
				case *net.IPAddr:
					ip = addr.(*net.IPAddr).IP
				}
				ipaddr := ip.To4().String()
				if ipaddr != "" && ipaddr != "<nil>" {
					return ipaddr
				}
			}
		}
	}
	return eth
}
