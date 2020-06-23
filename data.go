package main

import (
	"fmt"
	"net"
	"time"
)

func GetTimeString() string {
	t := time.Now()

	formatted := fmt.Sprintf("%s %s %d %02d:%02d:%02d %d",
		t.Weekday().String()[:3], t.Month().String()[:3], t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Year())
	return formatted
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Unknown"
}

func GetDataString(sourceName string) string {
	if sourceName == "TIME" {
		return GetTimeString()
	} else if sourceName == "IPADDRESS" {
		return GetLocalIP()
	} else {
		return "No data source for " + sourceName
	}
}
