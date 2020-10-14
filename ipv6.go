package main

import (
	"fmt"
	"log"
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

func main() {
	http.HandleFunc("/", getIPV6)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
