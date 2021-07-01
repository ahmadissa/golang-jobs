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
}
