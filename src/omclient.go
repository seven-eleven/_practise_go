package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var inputString string
var err error

func main() {
	common.ConfigInit()

	for {
		fmt.Println("Please enter a commond: <type 'h' for help; end with ';'>")

		inputReader = bufio.NewReader(os.Stdin)
		inputString, err = inputReader.ReadString(';')
		cmd := inputString[:len(inputString)-1]
		if nil == err {
			switch cmd {
			case "h":
				help()
			case "q":
				fmt.Println("quit")
				return
			case "s":
				stopServers()
			case "d":
				displayServerData()
			default:
				fmt.Println("Cmd not surpported.")
			}
		}
	}
}

func help() {
	info := `
		h -- help: get surpport command
		q -- quit: quit om client
		s -- stop: stop all the servers
		d -- display: display server statstistics for 10 times
	`

	fmt.Println(info)
}

func stopServers() {

}

func displayServerData() {
	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()
}
