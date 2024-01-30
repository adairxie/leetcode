package singleton

import (
	"fmt"
	"sync"
)

type single struct {
}

var lock sync.Mutex
var singleInstance *single

func getInstance() *single {
	if singleInstance != nil {
		fmt.Println("Single instance already created.")
		return singleInstance
	}

	lock.Lock()
	defer lock.Unlock()

	if singleInstance != nil {
		fmt.Println("Single instance already created.")
		return singleInstance
	}
	fmt.Println("Creating single instance now.")
	singleInstance = &single{}
	return singleInstance
}
