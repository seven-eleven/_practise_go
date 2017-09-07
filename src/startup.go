package main

import (
	"common"
	"os"
	"os/exec"
)

func main() {
	// get config from conf.json
	common.ConfigInit()
	server_ips := common.ConfigGetServerIpList()
	server_ports := common.ConfigGetServerPortList()

	// run servers
	for i := 0; i < len(server_ports); i++ {
		ip := server_ips[0] // only use local host ip
		port := server_ports[i]
		cmd := exec.Command("server/server.exe", ip, port)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if nil != err {
			common.Log(err.Error())
			return
		}
		common.Log("Server: ", ip, port, " Running")
	}

	// run clients
	cmd := exec.Command("client/client.exe")
	err := cmd.Run()
	if nil != err {
		common.Log(err.Error())
	}
	common.Log("A Client Running.")
}
