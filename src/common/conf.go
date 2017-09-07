package common

import (
	"encoding/json"
	"os"
	"strings"
)

type Config struct {
	server_ip_list      string
	server_port_list    string
	client_num_per_host int
	log_switch          bool
}

var configuration Config = Config{}

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
	if configuration == (Config{}) {
		return []string{}
	}

	ipLists := configuration.server_ip_list
	return strings.Split(ipLists, ";")
}

func ConfigGetServerPortList() []string {
	if configuration == (Config{}) {
		return []string{}
	}

	portLists := configuration.server_port_list
	return strings.Split(portLists, ";")
}

func ConfigGetClientNumPerHost() int {
	if configuration == (Config{}) {
		return 1
	}

	return configuration.client_num_per_host
}

func ConfigGetLogSwitch() bool {
	if configuration == (Config{}) {
		return true
	}

	return configuration.log_switch
}
