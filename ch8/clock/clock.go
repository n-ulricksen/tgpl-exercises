package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var location *time.Location
var port *int

func main() {
	// Get command line flags (port number)
	port = flag.Int("port", 8000, "port - default is 8000")
	flag.Parse()

	getTimeZone()

	listener := listenOnLocalhost()

	// Allow connections from multiple clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			// retry
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func listenOnLocalhost() net.Listener {
	host := "localhost:" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nListening on port %v\n", *port)

	return listener
}

func getTimeZone() {
	timezone := os.Getenv("TZ")
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal("Invalid location")
	}
	location = loc
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(location).
			Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
