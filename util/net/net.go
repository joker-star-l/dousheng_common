package util_net

import (
	"github.com/joker-star-l/dousheng_common/config/log"
	"net"
	"strconv"
)

func IpFromStr(ip string, port int) net.Addr {
	addr, err := net.ResolveTCPAddr("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		log.Slog.Errorln(err)
		return nil
	}
	return addr
}

func LocalIp() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		log.Slog.Errorln(err)
		return ""
	}
	for _, addr := range addr {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}
	log.Slog.Errorln("not found ipv4 address")
	return ""
}
