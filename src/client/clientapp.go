package client

import (
	"common"
	"net"
	"strconv"
)

// update key - value data to server
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
	msg := key + ";" + value
	err = common.SendMsg(conn, msg)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	err = common.RecvMsg(conn)
	if nil != err {
		common.Log(err.Error())
		return err
	}

	return nil
}

// send a stop command to server
func ClientHandleStop(conn net.Conn) {

}

// query key - value info from server
func ClientHandleQuery(conn net.Conn) {

}
