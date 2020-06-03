package errdemo

import (
	"errors"
	"fmt"
	"xiaobai.com/go_example/errdemo/sub/sub2"
)

func Diff(foo,bar int) error {
	if foo < 0 {
		return  errors.New("diff error")
	}
	if err:= sub2.Diff(foo,bar);err != nil {
		return  fmt.Errorf("sub1 error")
	}
	return nil
}

func IoDiff(foo,bar int) error  {
	_, err := sub2.IoDiff(foo, bar)
	return err
}