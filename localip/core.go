package localip

import (
	"log"
	"os/exec"
	"regexp"
)

// Ipv6 返回本地分配的 ipv6, 可能多个值
func Ipv6(ipv6Type string) []string {
	var reMatch string
	if ipv6Type == "fixed" {
		reMatch = `inet6 ([^f:][\da-f:]+)/\d+ scope global.?\n.+?valid_lft forever preferred_lft forever`
	} else {
		reMatch = `inet6 ([^f:][\da-f:]+)/\d+ scope global mngtmpaddr dynamic`
	}

	command := "ip addr show|grep -v deprecated|grep -A1 'inet6 [^f:]'"
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		log.Println("command err:", err)
		return []string{}
	}
	resp := string(bytes)

	re := regexp.MustCompile(reMatch)
	if re == nil {
		log.Println("regexp err")
		return []string{}
	}

	result := re.FindStringSubmatch(resp)

	if len(result) == 0 {
		log.Println("match err")
	}

	return result
}
