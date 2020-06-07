package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type Clock struct {
	City string
	Host string
	Time []byte
}

const DELIM = "="

var clocks []*Clock

func main() {
	if len(os.Args) <= 1 {
		printUsageError()
	}
	args := os.Args[1:]
	createClocksFromArgs(args)

	// Start the clocks - this will update Time fields on each server response
	for _, clock := range clocks {
		go startClock(clock)
	}

	// Print clock titles
	clockTitles := ""
	for _, clock := range clocks {
		clockTitles += clock.City
		clockTitles += "\t\t"
	}
	fmt.Println(clockTitles)

	// Repeatedly print clock times
	for {
		clockValues := ""
		for _, clock := range clocks {
			clockValues += string(clock.Time)
			clockValues += "\t"
		}
		fmt.Printf("%s\r", clockValues)
		time.Sleep(time.Second)
	}
}

func createClocksFromArgs(args []string) {
	for _, arg := range args {
		if !strings.Contains(arg, DELIM) {
			printUsageError()
		}

		parsed := strings.Split(arg, DELIM)
		if len(parsed) != 2 {
			printUsageError()
		}
		city := parsed[0]
		hostAddr := parsed[1]

		clocks = append(clocks, &Clock{
			city,
			hostAddr,
			[]byte{},
		})
	}
}

func startClock(clock *Clock) {
	conn, err := net.Dial("tcp", clock.Host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Read and store the time from the server to each Clock struct
	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		clock.Time = line[:len(line)-1]
	}
}

func printUsageError() {
	log.Fatal("Usage: ./main.go <city>=<host>:<port>")
}
