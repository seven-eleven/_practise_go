package main

import (
	"common"
	"os"
	"os/exec"
)

func main() {
	// run servers
	ip := "172.0.0.1"
	port := "23877"
	cmd := exec.Command("server/server.exe", ip, port)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if nil != err {
		common.Log(err.Error())
		return
	}
	common.Log("Server: ", ip, port, " Running")

	// run clients
	cmd = exec.Command("client/client.exe")
	err = cmd.Run()
	if nil != err {
		common.Log(err.Error())
	}
	common.Log("A Client Running.")
}
