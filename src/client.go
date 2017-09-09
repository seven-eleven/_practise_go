/*
   客户端程序
   每个CLIENT向所有SERVER进程建立连接，并不停发送key-value更新操作
*/
package main

import (
	"client"
	"common"
)

func main() {
	// 初始化配置
	err := common.ConfigInit()
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}

	common.Log("I am a client.")

	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()

	// 与所有SERVER进程建立连接并发起数据交互
	thread := 0
	for i := 0; i < len(serverIps); i++ {
		for j := 0; j < len(serverPorts); j++ {
			server := serverIps[i] + ":" + serverPorts[j]

			client.SetThreadState(thread, client.Alive)
			go client.ConnectToOneServer(server, thread, client.ClientHandleUpdate, true)
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
