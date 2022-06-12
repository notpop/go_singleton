package designPettern

import (
	"fmt"
	"time"
)

type singleton struct {
	time time.Time
}

var instance *singleton = newInstance()

func newInstance() *singleton {
	return &singleton{
		time: time.Now(),
	}
}

func getTime() time.Time {
	return instance.time
}

func GetInstance() *singleton {
	fmt.Printf("get instance at %s\n", getTime())
	return instance
}
