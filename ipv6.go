package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/uees/ipv6/local"
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
	ipv6, err := local.Ipv6(ipv6Type)
	if err != nil {
		fmt.Fprintf(writer, "ERROR: %s", err)
	}

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

	var addr string
	flag.StringVar(&addr, "addr", ":9090", "http listen address")
	flag.Parse()

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
