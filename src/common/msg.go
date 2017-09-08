package common

import (
	"net"
)

func RecvMsg(conn net.Conn) error {
	buffer := make([]byte, 2048)

	n, err := conn.Read(buffer)
	if nil != err {
		Log("Read msg error: ", err.Error())
		return err
	}

	// do something with the msg
	logHead := conn.LocalAddr().String() + " Recv From " + conn.RemoteAddr().String()
	Log(logHead, " Content: ", string(buffer[:n]))

	return nil
}

func SendMsg(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg))
	if nil != err {
		Log("Send msg error: ", err.Error())
		return err
	}

	logHead := conn.LocalAddr().String() + " Send To " + conn.RemoteAddr().String()
	Log(logHead, " Content: ", msg)

	return nil
}
