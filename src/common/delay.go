/*
   延时处理
*/

package common

import (
	"time"
)

// delay 1us
func Delay1US() {
	time.Sleep(time.Microsecond)
}

// delay 1ms = 1000us
func Delay1MS() {
	time.Sleep(time.Millisecond)
}

// delay n ms
func DelayNMS(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// delay 1s = 1000ms
func Delay1S() {
	time.Sleep(time.Second)
}

// delay n seconds
func DelayNS(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

// delay 1min
func Delay1Min() {
	time.Sleep(time.Minute)
}

// delay n minutes
func DelayNMin(n int) {
	time.Sleep(time.Duration(n) * time.Minute)
}
