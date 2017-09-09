/*
   配置信息管理，提供API
*/

package common

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerIps       string `json:"server_ip_list"`      // SERVER进程IP列表
	ServerPorts     string `json:"server_port_list"`    // SERVER进程PORT列表
	ServerApp       string `json:"server_app_name"`     // SERVER进程可执行文件名
	ClientNum       int    `json:"client_num_per_host"` // 单HOST上CLIENT数量
	ClientApp       string `json:"client_app_name"`     // CLIENT进程可执行文件名
	LogSwitch       bool   `json:"log_switch"`          // 日志打印开关
	UpdateIntervals int    `json:"update_intervals"`    // CLIENT key-value更新操作间隔，以ms为单位
}

var configuration = Config{}

func ConfigInit() error {
	file, err := os.Open("conf.json")
	if nil != err {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if nil != err {
		return err
	}

	return nil
}

func ConfigGetServerIpList() []string {
	if configuration == (Config{}) {
		return []string{}
	}

	ipLists := configuration.ServerIps

	return strings.Split(ipLists, ";")
}

func ConfigGetServerPortList() []string {
	if configuration == (Config{}) {
		return []string{}
	}

	portLists := configuration.ServerPorts

	return strings.Split(portLists, ";")
}

func ConfigGetClientNumPerHost() int {
	if configuration == (Config{}) {
		return 0
	}

	return configuration.ClientNum
}

func ConfigGetLogSwitch() bool {
	if configuration == (Config{}) {
		return true
	}

	return configuration.LogSwitch
}

func ConfigGetServerAppName() string {
	if configuration == (Config{}) {
		return ""
	}

	return configuration.ServerApp
}

func ConfigGetClientAppName() string {
	if configuration == (Config{}) {
		return ""
	}

	return configuration.ClientApp
}

func ConfigGetUpdateIntervals() int {
	if configuration == (Config{}) {
		return 1000
	}

	return configuration.UpdateIntervals
}

// 将所有配置信息以字符串的形式显示
func ConfigString() string {
	str := "Configuration:\r\n"
	str += "ServerIps: " + configuration.ServerIps + "\r\n"
	str += "ServerPorts: " + configuration.ServerPorts + "\r\n"
	str += "ServerApp: " + configuration.ServerApp + "\r\n"
	str += "ClientNum: " + strconv.Itoa(configuration.ClientNum) + "\r\n"
	str += "ClientApp: " + configuration.ClientApp + "\r\n"
	str += "LogSwitch: " + strconv.FormatBool(configuration.LogSwitch) + "\r\n"
	str += "UpdateIntervals: " + strconv.Itoa(configuration.UpdateIntervals) + "ms \r\n"

	return str
}
