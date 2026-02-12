package ui

import "fmt"

func PrintBanner() {
	fmt.Println(ColorCyan + `
██████╗ ███████╗ ██████╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝██╔════╝██╔═══██╗████╗  ██║
██████╔╝█████╗  ██║     ██║   ██║██╔██╗ ██║
██╔══██╗██╔══╝  ██║     ██║   ██║██║╚██╗██║
██║  ██║███████╗╚██████╗╚██████╔╝██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═══╝
` + ColorReset)
	// fmt.Println()

	fmt.Println("  " + Green("host") + " <ip>            - Host status (ping + nmap -sn)")
	fmt.Println("  " + Green("port") + " <ip>            - Port checker (nmap)")
	fmt.Println("  " + Cyan("help") + " / ?             - Show commands")
	fmt.Println("  " + Cyan("clear") + " cls            - Clear screen")
	fmt.Println("  " + Cyan("exit") + " / q             - Quit")
	fmt.Println()
}



func PrintBannerHelp() {
	fmt.Println(Yellow(`\\\ Recon - Help \\\`))
	fmt.Println(`recon check - used to ....`)

}