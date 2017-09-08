package main

import (
	"common"
	"os/exec"
)

func main() {
	// get config from conf.json
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

	// run servers
	for i := 0; i < len(serverPorts); i++ {
		ip := serverIps[0] // only use local host ip
		port := serverPorts[i]

		go runServer(serverApp, ip, port)
	}

	common.Delay1S() // wait servers to be built

	// run clients
	for i := 0; i < clientNum; i++ {
		go runClient(clientApp)
	}

	common.Delay1Min() // wait clients to be built
}

// run server app with os.exec
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

// run client app with os.exec
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
