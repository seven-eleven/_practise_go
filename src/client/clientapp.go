package client

import (
	"common"
	"net"
	"strconv"
)

// 构造update消息
func buildUpdateMsg(key string, value string) string {
	return common.CmdUpdate + ";" + key + ";" + value
}

// update key-value操作
func ClientHandleKeyValue(conn net.Conn) error {
	key := conn.LocalAddr().String() // ip:port

	value, err := common.KernelQuery(key)
	if nil != err {
		value = strconv.Itoa(1)
	} else {
		intValue, _ := strconv.Atoi(value)
		intValue += 1
		value = strconv.Itoa(intValue)
	}

	// update local
	common.KernelUpdate(key, value)

	// update to server
	msg := buildUpdateMsg(key, value)
	err = common.SendMsg(conn, msg)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	msg, err = common.RecvMsg(conn)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	return nil
}

// 构造stop消息
func buildStopMsg() string {
	return common.CmdStop + ";"
}

// send a stop command to server
func ClientHandleStop(conn net.Conn) {

}

// 构造query消息
func buildQueryMsg(key string) string {
	return common.CmdQuery + ";" + key
}

// query key - value info from server
func ClientHandleQuery(conn net.Conn) {

}
