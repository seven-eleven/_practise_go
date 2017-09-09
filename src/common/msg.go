/*
   消息收发处理
*/

package common

import (
	"net"
)

// 从连接端口读取消息
func RecvMsg(conn net.Conn) (string, error) {
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if nil != err {
		Log("Read msg error: ", err.Error())
		return "", err
	}

	// log for debug
	logHead := conn.LocalAddr().String() + " Recv From " + conn.RemoteAddr().String()
	Log(logHead, " Content: ", string(buffer[:n]))

	return string(buffer[:n]), nil
}

// 发送消息
func SendMsg(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg))
	if nil != err {
		Log("Send msg error: ", err.Error())
		return err
	}

	// log for debug
	logHead := conn.LocalAddr().String() + " Send To " + conn.RemoteAddr().String()
	Log(logHead, " Content: ", msg)

	return nil
}
