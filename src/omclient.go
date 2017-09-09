/*
   操作维护客户端
   支撑的命令可在程序运行后，输入"h;"获取帮助
*/

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
			case "c":
				displayConfiguration()
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
		c -- config: display configuration
	`

	fmt.Println(info)
}

// 通知SERVER停止服务
func stopServers() {
	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起查询
	thread := 0
	for i := 0; i < len(serverIps); i++ {
		for j := 0; j < len(serverPorts); j++ {
			server := serverIps[i] + ":" + serverPorts[j]

			go client.ConnectToOneServer(server, thread, client.ClientHandleStop, false)
			thread++
		}
	}
}

// 显示SERVER数据
func displayServerData() {
	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起查询
	thread := 0
	for i := 0; i < len(serverIps); i++ {
		for j := 0; j < len(serverPorts); j++ {
			server := serverIps[i] + ":" + serverPorts[j]

			client.SetThreadState(thread, client.Alive)
			go client.ConnectToOneServer(server, thread, client.ClientHandleQuery, false)
			thread++
		}
	}

	// 等待所有子线程退出后主线程退出
	for {
		common.Delay1S()

		mainThreadOver := true
		for i := 0; i < thread; i++ {
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

// 显示配置信息
func displayConfiguration() {
	fmt.Println(common.ConfigString())
}
