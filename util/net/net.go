package util_net

import (
	"github.com/joker-star-l/dousheng_common/config/log"
	"net"
	"os"
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
	host, _ := os.Hostname()
	addr, _ := net.LookupIP(host)
	for _, addr := range addr {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	log.Slog.Errorln("not found ipv4 address")
	return "127.0.0.1"
}
