package main

import (
	b64 "encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"

	"code.cloudfoundry.org/bytefmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/mackerelio/go-osstat/memory"
	"github.com/matishsiao/goInfo"
	"github.com/reiver/go-telnet"
	"github.com/shirou/gopsutil/load"
)

func main() {
	args := os.Args
	argslen := len(args)

	if argslen == 3 {
		// fmt.Println("[@command]", args[0], args[1], args[2])

		switch args[1] {

		case "b64e":
			enc := b64.StdEncoding.EncodeToString([]byte(args[2]))
			// fmt.Println("[@Base64 Encoding] input:", args[2], "==> output: ", enc)
			fmt.Println(enc)

		case "b64d":
			dec, _ := b64.StdEncoding.DecodeString(args[2])
			// fmt.Println("[@Base64 decoding] input:", args[2], "==> output:", string(dec))
			fmt.Println(string(dec))

		case "tel":

			if strings.Contains(args[2], ":") {
				telArgs := strings.Split(args[2], ":")

				if !valid.IsInt(telArgs[1]) {
					fmt.Println("[@command]", args[0], args[1], args[2])
					fmt.Println("[@error] port must be number")
				}

				var caller telnet.Caller = telnet.StandardCaller
				telnet.DialToAndCall(args[2], caller)

			} else {
				fmt.Println("[@command]", args[0], args[1], args[2])
				fmt.Println("[@error] argument must in the form \"host:port\".")
			}

		default:
			fmt.Println("[@command]", args[0], args[1], args[2])
			fmt.Println("[@error]", args[1], "is unknown <command>")
			fmt.Println("Run 'jtl help' for usage.")
			fmt.Println()
		}

	} else if argslen == 2 {

		switch args[1] {

		case "help":
			fmt.Println(
				`jtl is a collection of tools that can be used in the cli environment. (developed by pjw)

Usage:
    jtl <command> [arguments]

           b64e      Outputs the Base64 encoded result of the arguments.
           b64d      Outputs the Base64 decoded result of the arguments.
           tel       This is a telnet client. argument must in the form "host:port".
                     Only supported PLAIN communication, SSL, TLS communication is not supported.
                     When you exit the shell "Ctrl + c" or after the connection is closed, press "Enter twice".
           sys       Outputs the OS platform Information. [ Kernel, Core, Platform, OS, Hostname, CPUs, Loadavg, Ip(Only Local IPV4), Memory ]
                     No require arguments.



										 `)
		case "version":
			fmt.Println("jtl version \"0.0.2\"")

		case "v":
			fmt.Println("jtl version \"0.0.2\"")

		case "sys":
			// fmt.Println("[@command]", args[0], args[1])
			gi := goInfo.GetInfo()
			gi.VarDump()
			fmt.Println("Ip (Local):", GetLocalIP())

			memory, err := memory.Get()

			if err != nil {
				fmt.Println("[@command]", args[0], args[1])
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return
			}

			load, _ := load.Avg()
			fmt.Println("Loadavg (1m, 5m, 15m):", fmt.Sprintf("%.2f", load.Load1), fmt.Sprintf("%.2f", load.Load5), fmt.Sprintf("%.2f", load.Load15))

			fmt.Printf("Memory (Total): %d MB\n", memory.Total/bytefmt.MEGABYTE)
			fmt.Printf("Memory (Used): %d MB\n", memory.Used/bytefmt.MEGABYTE)
			fmt.Printf("Memory (Cached): %d MB\n", memory.Cached/bytefmt.MEGABYTE)
			fmt.Printf("Memory (Free): %d MB\n", memory.Free/bytefmt.MEGABYTE)

		default:
			fmt.Println("[@command]", args[0], args[1])
			fmt.Println("[@error] [arguments] does not exist.")
			fmt.Println("Run 'jtl help' for usage.")
		}
	} else if argslen == 1 {
		fmt.Println("[@error] <command> does not exist.")
		fmt.Println("Run 'jtl help' for usage.")
	}

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
	return ""
}
