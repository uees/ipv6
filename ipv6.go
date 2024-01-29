package main

import (
	"fmt"
	"log"
	"main/localip"
	"net/http"

	"github.com/xjh22222228/ip"
)

// getIPV6ByNet 利用网络 API 获取 IPV6
func getIPV6ByNet(writer http.ResponseWriter, request *http.Request) {
	ipv6, err := ip.V6()
	if err != nil {
		fmt.Fprintf(writer, "ERROR: %v", err)
	} else {
		fmt.Fprintf(writer, "%s", ipv6)
	}
}

// getIPV6ByShell 利用本地 Shell command 获取 IPV6
func getIPV6ByShell(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	ipv6Type := request.Form.Get("type")
	ipv6 := localip.Ipv6(ipv6Type)

	if len(ipv6) == 0 {
		fmt.Fprintf(writer, "ERROR: %s", "not get ipv6")
	} else if len(ipv6) == 1 {
		fmt.Fprintf(writer, "%s", ipv6[0])
	} else {
		fmt.Fprintf(writer, "%s", ipv6[1])
	}
}

func main() {
	http.HandleFunc("/ipv6-request-net", getIPV6ByNet)
	http.HandleFunc("/ipv6", getIPV6ByShell)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
