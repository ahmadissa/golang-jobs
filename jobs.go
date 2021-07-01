package Jobs

import (
	"errors"
	"sync"
)

var mutex = &sync.Mutex{}
var list = make(map[string]func(...interface{}) error)

//Add Job to jobs list
func Add(key string, job func(...interface{}) error) {
	mutex.Lock()
	defer mutex.Unlock()
	list[key] = job
}

//Run job from list
func Run(key string, attribute ...interface{}) error {
	job, ok := list[key]
	if !ok {
		return errors.New("Job not found:" + key)
	}
	return job(attribute...)
}
