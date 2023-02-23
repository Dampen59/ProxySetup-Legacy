package main

import (
	"fmt"
	"goproxy-master"
	"log"
	"net"
	"net/http"
)

var networkInterface = "192.168.0.0" // Change to your network interface IPv4. Windows => "ipconfig /all", Linux => "ifconfig" if you have net-tools installed or "ip a"
var listenPort = 8080 // Change to your convenience. Note that you will have to open this port in your Firewall / Router

func runProxyServer() {
	nInterface := fmt.Sprintf("%v:0", networkInterface)
	localAddr, err := net.ResolveTCPAddr("tcp", nInterface)

	if err != nil {
		panic(err)
	}

	d := net.Dialer{LocalAddr: localAddr}
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Tr.Dial = d.Dial

	proxyPort := fmt.Sprintf(":%v", listenPort)
	log.Println(fmt.Sprintf("Proxy sur l'interface %v en écoute sur le port %v", nInterface, proxyPort))
	log.Fatal(http.ListenAndServe(proxyPort, proxy))
}

func main() {
	runProxyServer()
}