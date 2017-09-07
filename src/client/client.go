package main

import (
	"common"
	"net"
	"strconv"
	"time"
)

func main() {
	server := "127.0.0.1:23877"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}
	common.Log("connect success")

	var times int
	for {
		msg := "Hello World! Times " + strconv.Itoa(times)
		err := common.SendMsg(conn, msg)
		if nil != err {
			break
		}

		err = common.RecvMsg(conn)
		if nil != err {
			break
		}

		time.Sleep(10000000)
		times++
	}
}
