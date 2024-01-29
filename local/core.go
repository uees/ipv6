package local

import (
	"errors"
	"log"
	"os/exec"
	"regexp"
)

// Ipv6 返回本地分配的 ipv6, 可能多个值
func Ipv6(ipv6Type string) ([]string, error) {
	var reMatch string

	if ipv6Type == "fixed" {
		reMatch = `inet6 ([^f:][\da-f:]+)/\d+ scope global.?\n.+?valid_lft forever preferred_lft forever`
	} else {
		reMatch = `inet6 ([^f:][\da-f:]+)/\d+ scope global dynamic`
	}

	command := "ip addr show|grep -v deprecated|grep -A1 'inet6 [^f:]'"
	cmd := exec.Command("/bin/sh", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		log.Println("command err:", err)
		return nil, err
	}
	resp := string(bytes)

	re := regexp.MustCompile(reMatch)

	result := re.FindAllStringSubmatch(resp, -1)

	if len(result) == 0 {
		return nil, errors.New("ipv6 match err")
	}

	ips := []string{}
	for _, item := range result {
		ips = append(ips, item[1])
	}

	return ips, nil
}
