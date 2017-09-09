/*
   支撑CLIENT与SERVER建立连接，并根据每次消息处理后的反馈判断连接是否继续保留
*/

package client

import (
	"common"
	"net"
)

const (
	Alive = true
	Dead  = false
)

/* CLIENT的连接子线程状态
   每个CLIENT要跟所有的SERVER建立连接，因此子线程最大个数为SERVER进程最大个数
   维护这个信息是为了让主线程判断是否可以退出
*/
var threadState [common.MaxServer]bool

func SetThreadState(thread int, state bool) {
	threadState[thread] = state
}

func GetThreadState(thread int) bool {
	return threadState[thread]
}

/* CLIENT子线程内的连接处理函数
   input: SERVER地址， 线程ID， 消息发送函数， 是否循环执行
*/
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

	// 根据交互间隔，反复处理消息发送和接受
	intervals := common.ConfigGetUpdateIntervals()
	for {
		err = f(conn)
		if nil != err {
			common.Log(err.Error())
			break
		}

		if !loop {
			break
		}

		common.DelayNMS(intervals)
	}
}
