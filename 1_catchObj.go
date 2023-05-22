package try

import (
	"errors"
	"fmt"
)

type catchObj struct {
	logic func()
}

func (obj *catchObj) Catch(cb func(info interface{})) {
	defer func() {
		errInfo := recover()
		if errInfo != nil {
			if cb != nil {
				cb(errInfo)
			}
		}
	}()
	if obj.logic != nil {
		obj.logic()
	}
}

func (obj *catchObj) Err() (err error) {
	obj.Catch(func(panicInfo interface{}) {
		err = obj.panicInfoToErr(panicInfo)
	})
	return err
}

func (obj *catchObj) ErrMsg() (errMsg string) {
	obj.Catch(func(panicInfo interface{}) {
		errMsg = obj.panicInfoToString(panicInfo)
	})
	return errMsg
}

func (obj *catchObj) Stack() (stack string) {
	obj.Catch(func(interface{}) {
		stack = readStack()
	})
	return
}

func (obj *catchObj) ErrMsgAndStack() (errMsg string, stack string) {
	obj.Catch(func(panicInfo interface{}) {
		errMsg = obj.panicInfoToString(panicInfo)
		stack = readStack()
	})
	return
}

func (obj *catchObj) ErrAndStack() (err error, stack string) {
	obj.Catch(func(panicInfo interface{}) {
		err = obj.panicInfoToErr(panicInfo)
		stack = readStack()
	})
	return
}

func (obj *catchObj) panicInfoToString(panicInfo interface{}) string {
	return fmt.Sprintf("%v", panicInfo)
}

func (obj *catchObj) panicInfoToErr(panicInfo interface{}) error {
	return errors.New(obj.panicInfoToString(panicInfo))
}
