// client.go
package main

import (
	"client"
	"common"
)

func main() {
	err := common.ConfigInit()
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}

	common.Log("I am a client.")

	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起数据交互
	for i := 0; i < len(serverPorts); i++ {
		server := serverIps[0] + ":" + serverPorts[i]

		client.SetThreadState(i, client.Alive)
		go client.ConnectToOneServer(server, i, client.ClientHandleUpdate, true)
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
