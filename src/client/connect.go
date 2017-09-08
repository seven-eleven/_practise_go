package client

import (
	"common"
	"net"
)

const (
	Alive = true
	Dead  = false
)

var threadState [common.MaxServer]bool // flags for thread alive or dead

func SetThreadState(thread int, state bool) {
	threadState[thread] = state
}

func GetThreadState(thread int) bool {
	return threadState[thread]
}

// client connect & communicate to one server
func ConnectToOneServer(server string, thread int, f func(net.Conn) error, loop bool) {
	defer SetThreadState(thread, Dead) // set flag when thread over

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
	common.Log("Connect to server success, server: ", conn.RemoteAddr().String())

	// handle msg
	for {
		err = f(conn)
		if nil != err {
			common.Log(err.Error())
			break
		}

		if !loop {
			break
		}

		common.Delay1S()
	}
}
