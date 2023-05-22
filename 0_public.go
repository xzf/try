package try

import (
	"bytes"
	"errors"
	"fmt"
	"runtime"
)

type catchObj struct {
	logic func()
}

func (obj *catchObj) Catch(cb func(info interface{})) (err error) {
	defer func() {
		errInfo := recover()
		if errInfo != nil {
			err = errors.New(obj.panicInfoToString(errInfo))
			if cb != nil {
				cb(errInfo)
			}
		}
	}()
	if obj.logic != nil {
		obj.logic()
	}
	return nil
}

func (obj *catchObj) Error() (err error) {
	obj.Catch(func(panicInfo interface{}) {
		err = errors.New(obj.panicInfoToString(panicInfo))
	})
	return err
}

func (obj *catchObj) ErrMsg() (errMsg string) {
	obj.Catch(func(panicInfo interface{}) {
		errMsg = obj.panicInfoToString(panicInfo)
	})
	return errMsg
}

func (obj *catchObj) panicInfoToString(panicInfo interface{}) string {
	return fmt.Sprintf("%v", panicInfo)
}

func (obj *catchObj) Stack() (stack string) {
	obj.Catch(func(interface{}) {
		stack = readStack()
	})
	return
}

func Try(logic func()) *catchObj {
	return &catchObj{
		logic: logic,
	}
}

func readStack() string {
	content := make([]byte, 102400)
	index := runtime.Stack(content, true)
	if index > 0 {
		end := bytes.Index(content[1:], []byte(`
goroutine `))
		if end > 0 {
			content = content[:end]
		}
	}
	return string(content)
}
