package designPettern

import (
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

func GetInstance() (*singleton, string) {
	return instance, "get instance at" + getTime().String()
}
