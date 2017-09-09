/*
   SERVER进程运行程序
   如果侦听到连接，就会创建线程处理当前连接相关的消息收发
*/

package main

import (
	"common"
	"net"
	"os"
	"server"
)

func main() {
	// 读取配置
	err := common.ConfigInit()
	if nil != err {
		common.Log("Fatal error: ", err.Error())
		return
	}

	// 传入参数: appname, serverip, serverport
	args := os.Args
	argc := len(os.Args)
	if 3 != argc {
		common.Log("Args num for server error.")
		return
	}
	ip := args[1]
	port := args[2]
	server := ip + ":" + port

	common.Log("I am a server: ", server)

	netListen, err := net.Listen("tcp", server)
	if nil != err {
		common.Log("Tcp listen error.")
		return
	}

	defer netListen.Close()

	// 侦听连接并启动处理线程
	var conn net.Conn
	for {
		common.Log("Waiting for clients")
		conn, err = netListen.Accept()
		if nil != err {
			continue
		}

		common.Log(conn.RemoteAddr().String(), " tcp connect success.")

		go serverHandle(conn)
	}
}

// SERVER线程处理函数
func serverHandle(conn net.Conn) {
	for {
		msg, err := common.RecvMsg(conn)
		if nil != err {
			return
		}

		loop := server.ServerAppMain(conn, msg)
		if !loop {
			break
		}
	}

	common.Log("Server handle thread over.")
}
