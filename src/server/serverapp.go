package server

import (
	"common"
	"net"
	"strings"
)

/* SERVER处理线程内消息处理
   input: 连接, 消息
   output: 是否继续接收消息处理
*/
func ServerAppMain(conn net.Conn, msg string) bool {
	cmd, content := decodeServerMsg(msg)

	var loop bool
	switch cmd {
	case common.CmdUpdate:
		serverHandleUpdate(conn, content)
		loop = true
	case common.CmdQuery:
		serverHandleQuery(conn, content)
		loop = false
	case common.CmdStop:
		serverHandleStop()
		loop = false
	default:
		common.Log("invalid command at server")
		loop = false
	}

	return loop
}

/* 解析Client发送过来的消息
   input: 消息
   output: 操作类型, 消息内容
*/
func decodeServerMsg(msg string) (string, string) {
	index := strings.Index(msg, ";")
	if -1 == index {
		return "", ""
	}

	cmd := msg[:index]
	content := msg[index+1:]

	return cmd, content
}

// SERVER处理update key-value操作
func serverHandleUpdate(conn net.Conn, content string) {
	pieces := strings.Split(content, ";")
	if 2 != len(pieces) {
		common.Log("content err for update")
		return
	}

	key := pieces[0]
	value := pieces[1]

	// 本地更新
	common.KernelUpdate(key, value)

	// 回复确认消息
	ack := buildAckMsg()
	err := common.SendMsg(conn, ack)
	if nil != err {
		common.Log(err.Error())
		return
	}
}

// SERVER处理query操作
func serverHandleQuery(conn net.Conn, content string) {

}

// SERVER处理stop操作
func serverHandleStop() {

}

// 构造ack消息
func buildAckMsg() string {
	return common.CmdAck + ";"
}
