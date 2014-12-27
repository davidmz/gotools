package mustbe_test

import (
	"errors"
	"testing"

	"github.com/davidmz/gotools/mustbe"
)

var (
	testError = errors.New("Test Error")
)

func TestCatch(t *testing.T) {
	i := new(int)
	*i = 1
	if foo, err := testCatch(returnsError); foo != nil || err != testError {
		t.Errorf("testCatch returns %v, %v, want %v, %v", foo, err, nil, testError)
	}
	if foo, err := testCatch(returnsOK); foo == nil || *foo != 1 || err != nil {
		t.Errorf("testCatch returns *%v, %v, want *%v, %v", *foo, err, *i, nil)
	}
}

func testCatch(fun func() error) (foo *int, err error) {
	defer mustbe.Catched(&err, &foo)
	foo = new(int)
	*foo = 1
	mustbe.OK(fun())
	return
}

func returnsError() error {
	return testError
}

func returnsOK() error {
	return nil
}
