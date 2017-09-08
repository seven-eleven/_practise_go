package common

import (
	"errors"
	"sync"
)

var KernelMap = make(map[string]string)

var rwlock = new(sync.RWMutex)

// update data by (key, value)
func KernelUpdate(key string, value string) {
	if "" == key {
		Log("key is null")
		return
	}

	rwlock.Lock()
	KernelMap[key] = value
	Log("kernel write: ", key, ",", value)
	rwlock.Unlock()
}

// query value by key
func KernelQuery(key string) (string, error) {
	if "" == key {
		Log("key is null")
		return "", errors.New("key is null")
	}

	rwlock.RLock()
	value, ok := KernelMap[key]
	Log("kernel read by ", key)
	rwlock.RUnlock()

	if ok {
		return value, nil
	} else {
		return "", errors.New("key not found")
	}
}
