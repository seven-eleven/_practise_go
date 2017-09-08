package common

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	ServerIps   string `json:"server_ip_list"`
	ServerPorts string `json:"server_port_list"`
	ServerApp   string `json:"server_app_name"`
	ClientNum   int    `json:"client_num_per_host"`
	ClientApp   string `json:"client_app_name"`
	LogSwitch   bool   `json:"log_switch"`
}

var configuration = Config{}

func ConfigInit() error {
	file, err := os.Open("conf.json")
	if nil != err {
		Log("Open config json file fail.")
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if nil != err {
		Log("Decode config fail, ", err.Error())
		return err
	}

	return nil
}

func ConfigGetServerIpList() []string {
	ipLists := configuration.ServerIps
	Log("Configuration: server ip list ", ipLists)

	return strings.Split(ipLists, ";")
}

func ConfigGetServerPortList() []string {
	portLists := configuration.ServerPorts
	Log("Configuration: server port list ", portLists)

	return strings.Split(portLists, ";")
}

func ConfigGetClientNumPerHost() int {
	Log("Configuration: client num ", strconv.Itoa(configuration.ClientNum))

	return configuration.ClientNum
}

func ConfigGetLogSwitch() bool {
	Log("Configuration: log switch ", strconv.FormatBool(configuration.LogSwitch))

	return configuration.LogSwitch
}

func ConfigGetServerAppName() string {
	Log("Configuration: server app name ", configuration.ServerApp)

	return configuration.ServerApp
}

func ConfigGetClientAppName() string {
	Log("Configuration: client app name ", configuration.ClientApp)

	return configuration.ClientApp
}
