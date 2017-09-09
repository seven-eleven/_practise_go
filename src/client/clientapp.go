package client

import (
	"common"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// update key-value操作
func ClientHandleUpdate(conn net.Conn) error {
	key := conn.LocalAddr().String() // ip:port

	value, err := common.KernelQueryByKey(key)
	if nil != err {
		value = strconv.Itoa(1)
	} else {
		intValue, _ := strconv.Atoi(value)
		intValue += 1
		value = strconv.Itoa(intValue)
	}

	// 更新本地数据
	common.KernelUpdate(key, value)

	// 更新SERVER数据
	msg := buildUpdateMsg(key, value)
	err = common.SendMsg(conn, msg)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	// 接收ACK
	msg, err = common.RecvMsg(conn)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	return nil
}

// 向SERVER发起查询所有客户端访问次数
func ClientHandleQuery(conn net.Conn) error {
	// 向SERVER发起查询
	msg := buildQueryMsg()
	err := common.SendMsg(conn, msg)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	// 接收ACK
	msg, err = common.RecvMsg(conn)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	// 处理ACK
	clientHandleAck4Query(msg)

	return nil
}

// Client发送stop命令给SERVER
func ClientHandleStop(conn net.Conn) error {
	msg := buildStopMsg()
	err := common.SendMsg(conn, msg)
	if nil != err {
		common.Log(err.Error())
		return err
	}
	return nil
}

// Client处理ack4query消息
func clientHandleAck4Query(msg string) {
	index := strings.Index(msg, ";")

	//common.Log(msg[index+1:])
	fmt.Println(msg[index+1:])
}

// 构造update消息
func buildUpdateMsg(key string, value string) string {
	return common.CmdUpdate + ";" + key + ";" + value
}

// 构造query消息
func buildQueryMsg() string {
	return common.CmdQuery + ";"
}

// 构造stop消息
func buildStopMsg() string {
	return common.CmdStop + ";"
}
