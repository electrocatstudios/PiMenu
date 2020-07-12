package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
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

type Version struct {
	Major    int `json:"major"`
	Minor    int `json:"minor"`
	Revision int `json:"revision"`
}

func GetVersion() string {

	filename := fmt.Sprintf("version.json")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "Version: Unknown"
	}

	file, _ := ioutil.ReadFile(filename)
	var ver Version
	err := json.Unmarshal([]byte(file), &ver)
	if err != nil {
		return "Version: Error"
	}
	var ret string
	ret = fmt.Sprintf("Version: %i.%i.%i", ver.Major, ver.Minor, ver.Revision)
	return ret
}

func GetDataString(sourceName string) string {
	if sourceName == "TIME" {
		return GetTimeString()
	} else if sourceName == "IPADDRESS" {
		return GetLocalIP()
	} else if sourceName == "VERSION" {
		return GetVersion()
	} else {
		return "No data source for " + sourceName
	}
}
