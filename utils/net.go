package utils

import (
	"fmt"
	"net"
)

//GetAllIPv4Address 获取所有IP v4地址
func GetAllIPv4Address() []net.IP {
	result := []net.IP{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return result
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				result = append(result, ipnet.IP)
			}
		}
	}
	return result
}
