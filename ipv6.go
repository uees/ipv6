package main

import (
	"fmt"
	"log"
	"main/localip"
	"net/http"

	"github.com/xjh22222228/ip"
)

// 获取本机 IPV6
func getIPV6(writer http.ResponseWriter, request *http.Request) {
	ipv6, err := ip.V6()
	if err != nil {
		fmt.Fprintf(writer, "ERROR: %v", err)
	} else {
		fmt.Fprintf(writer, ipv6)
	}
}

// getIPV6Local 利用本地 API
func getIPV6Local(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	ipv6Type := request.Form.Get("type")
	ipv6 := localip.Ipv6(ipv6Type)

	if ipv6 == "" {
		fmt.Fprintf(writer, "ERROR: not get ipv6")
	} else {
		fmt.Fprintf(writer, ipv6)
	}
}

func main() {
	http.HandleFunc("/", getIPV6)
	http.HandleFunc("/localip", getIPV6Local)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
