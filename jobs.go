package Jobs

import (
	"errors"
	"sync"
)

var mutex = &sync.Mutex{}
var list = make(map[string]func(...interface{}) error)
var listStr = make(map[string]func(...string) error)

//Add Job to jobs list
func Add(key string, job func(...interface{}) error) {
	mutex.Lock()
	defer mutex.Unlock()
	list[key] = job
}

//AddStr Job with string args
func AddStr(key string, job func(...string) error) {
	mutex.Lock()
	defer mutex.Unlock()
	listStr[key] = job
}

//RunStr job from jobs list added by AddStr
func RunStr(key string, attribute ...string) error {
	job, ok := listStr[key]
	if !ok {
		return errors.New("Job not found:" + key)
	}
	return job(attribute...)
}

//Run job from list
func Run(key string, attribute ...interface{}) error {
	job, ok := list[key]
	if !ok {
		return errors.New("Job not found:" + key)
	}
	return job(attribute...)
}
