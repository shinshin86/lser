package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

func ip() (string, error) {
	var ip string
	netInterfaceAddresses, _ := net.InterfaceAddrs()

	for _, addrs := range netInterfaceAddresses {
		networkIp, ok := addrs.(*net.IPNet)

		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip = networkIp.IP.String()
		}
	}

	if ip == "" {
		return "", fmt.Errorf("Could not get IP address")
	}

	return ip, nil
}

func main() {
	var p = flag.Int("p", 3000, "port number")
	flag.Parse()

	port := ":" + strconv.Itoa(*p)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)
	fmt.Println("[lser] Start local server")
	fmt.Println("----- YOUR LOCAL URL -----")
	fmt.Printf("http://localhost:%d\n", *p)

	ip, err := ip()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("http://%s:%d\n", ip, *p)

	http.ListenAndServe(port, nil)
}
