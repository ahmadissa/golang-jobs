package Jobs

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	err := Run("noFunction")
	if err == nil {
		t.Error(errors.New("function does not exists"))
	}
	testFN := func(x ...interface{}) error {
		return nil
	}
	testFNStr := func(x ...string) error {
		return nil
	}
	Add("testFN", testFN)
	err = Run("testFN")
	if err != nil {
		t.Error(err)
	}
	Add("testFN1", testFN)
	err = Run("testFN1")
	if err != nil {
		t.Error(err)
	}
	AddStr("testFNStr", testFNStr)
	err = RunStr("testFNStr")
	if err != nil {
		t.Error(err)
	}
}
