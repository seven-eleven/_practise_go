package main

import (
	"bufio"
	"client"
	"common"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var inputString string
var err error

// 操作维护客户端入口
func main() {
	common.ConfigInit() // 初始化配置信息

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

// 打印帮助信息
func help() {
	info := `
		h -- help: get surpport command
		q -- quit: quit om client
		s -- stop: stop all the servers
		d -- display: display server statstistics for 10 times
	`

	fmt.Println(info)
}

// 通知SERVER停止服务
func stopServers() {
	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起查询
	for i := 0; i < len(serverPorts); i++ {
		server := serverIps[0] + ":" + serverPorts[i]

		go client.ConnectToOneServer(server, i, client.ClientHandleStop, false)
	}
}

// 显示SERVER数据
func displayServerData() {
	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起查询
	for i := 0; i < len(serverPorts); i++ {
		server := serverIps[0] + ":" + serverPorts[i]

		client.SetThreadState(i, client.Alive)
		go client.ConnectToOneServer(server, i, client.ClientHandleQuery, false)
	}

	// 等待所有子线程退出后主线程退出
	for {
		common.Delay1S()

		mainThreadOver := true
		for i := 0; i < len(serverPorts); i++ {
			if client.GetThreadState(i) {
				mainThreadOver = false
				break
			}
		}

		if mainThreadOver {
			break
		}
	}
}
