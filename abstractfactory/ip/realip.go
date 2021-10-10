package ip

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request)  {
	ip , err := getIp(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no valid ip"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(ip))
}
func getIp(r *http.Request) (string, error) {

	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	ips := r.Header.Get("X-FORWARDED-FOE")
	splitIPs := strings.Split(ips, ",")
	for _, ip := range splitIPs { //第一个IP
		netIp := net.ParseIP(ip)
		if netIp != nil {
			return ip, nil
		}
	}
	ip,_, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}