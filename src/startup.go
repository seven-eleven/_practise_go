/*
   启动系统
   先根据配置中SERVER IP列表，如果有属于本主机IP，会用配置文件中的SERVER端口列表，启动对应数据的SERVER进程
   再根据一个主机上的CLIENT数目，启动CLIENT
*/

package main

import (
	"common"
	"net"
	"os/exec"
)

func main() {
	// 读取配置信息
	err := common.ConfigInit()
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}

	serverIps := common.ConfigGetServerIpList()
	serverPorts := common.ConfigGetServerPortList()
	serverApp := common.ConfigGetServerAppName()
	clientNum := common.ConfigGetClientNumPerHost()
	clientApp := common.ConfigGetClientAppName()

	// 根据配置运行SERVER
	for i := 0; i < len(serverIps); i++ {
		ip := serverIps[i]
		local := isLocalIP(ip)
		if !local {
			continue
		}

		for j := 0; j < len(serverPorts); j++ {
			port := serverPorts[j]
			go runServer(serverApp, ip, port)
		}
	}

	common.Delay1S() // wait servers to be built

	// 根据配置的一个主机上的客户端数目启动CLIENT
	for i := 0; i < clientNum; i++ {
		go runClient(clientApp)
	}

	common.Delay1Min() // 等待SERVER & CLIENT进程启动完全后退出
}

// 启动一个SERVER进程
func runServer(app string, ip string, port string) {
	if "" == app || "" == ip || "" == port {
		common.Log("Input err for runServer.")
		return
	}

	common.Log("Run a server: ", ip, ":", port)

	cmd := exec.Command(app, ip, port)
	err := cmd.Run()
	if nil != err {
		common.Log(err.Error())
	}
}

// 启动一个CLIENT进程
func runClient(app string) {
	if "" == app {
		common.Log("App name is null for runClient.")
		return
	}

	common.Log("Run a client.")

	cmd := exec.Command(app)
	err := cmd.Run()
	if nil != err {
		common.Log(err.Error())
	}
}

// 判断IP是否为本机IP
func isLocalIP(ip string) bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false
	}
	for i := range addrs {
		intf, _, err := net.ParseCIDR(addrs[i].String())
		if err != nil {
			return false
		}
		if net.ParseIP(ip).Equal(intf) {
			return true
		}
	}
	return false
}
