package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"

	"github.com/ziutek/telnet"
)

func main() {

	// init an slice (containing all the open ports)
	var ports []int

	// open ports
	var openPorts = make(chan int, 32000)

	// Get hostname from cli (if given)
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			Usage()
			os.Exit(-1)
		} else {
			host = os.Args[1]
		}
	}

	// Loop through the ports and invoke telnet_dialer
	for port := 1; port < 32001; port++ {
		limit <- true
		go TelnetDialer(port, openPorts)
	}

	for i := 0; i < 32000; i++ {
		openPort := <-openPorts
		if openPort != 0 {
			ports = append(ports, openPort)
		}
	}
	sort.Ints(ports)
	fmt.Println("\nOpen ports: ", host, ":", ports)
}

// Telnet dialer
func TelnetDialer(port int, openPorts chan int) {
	fmt.Printf("\rScanning %s ports:%d/32000", host, port)
	_, err := telnet.DialTimeout("tcp", host+":"+strconv.Itoa(port), time.Second*1)
	if err == nil {
		openPorts <- port
	} else {
		openPorts <- 0
	}
	<-limit
}

// Poor man's usage
func Usage() {
	usage := `
	Usage:
	go run scan_open_ports.go [<hostname>|-h|--help]
	./scan_open_ports [<hostname>|-h|--help]

	E.g.
	./scan_open_ports # Scans localhost for all open ports

	./scan_open_ports server1.zyx.com # Scan host server1.zyx.com for all open ports

	   	`
	fmt.Println(usage)
}

// Limit concurrency based on the cpus
var limit = make(chan bool, runtime.NumCPU()*100) // 100 times number of cores

// default host if not given
var host = "localhost"
