package common

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var logFile string = "" // log file belongs to one server process
var logSwitch bool = true

func getTimeString(now time.Time) string {
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()

	retString := fmt.Sprintf("%d_%02d_%02d_%02d_%02d_%02d", year, int(mon), day, hour, min, sec)
	return retString
}

// init for logFile & logSwitch
func logInit() {
	if "" == logFile {
		now := time.Now()
		rand.Seed(now.UnixNano())
		random := rand.Intn(10000)

		// time + pid + rand

		logFile = getTimeString(now) + "_" + strconv.Itoa(os.Getpid()) + "_" + strconv.Itoa(random) + ".log"
	}

	logSwitch = true
}

// check if file exist
func checkFileIsExist(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}

// print log into logFile
func Log(v ...interface{}) {
	if !logSwitch {
		return
	}

	logInit()

	var fd *os.File
	var err error
	if checkFileIsExist(logFile) {
		fd, err = os.OpenFile(logFile, os.O_APPEND, 0666)
	} else {
		fd, err = os.Create(logFile)
	}
	if nil != err {
		fmt.Println(err.Error())
		return
	}

	content := getTimeString(time.Now()) + " "
	for _, str := range v {
		content += str.(string)
	}
	content += "\r\n"
	fd.Write([]byte(content))

	fd.Close()
}
