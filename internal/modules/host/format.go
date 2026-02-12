package host

import "fmt"

func PrintResult(r Result) {
	fmt.Printf("%s : %s\n", r.Target, r.Status)
	fmt.Printf("    Signal: PING=%s | NMAP_SN=%s\n", r.Ping, r.NmapSN)
	fmt.Printf("    Hint  : %s\n", r.Hint)
	fmt.Println()
}
