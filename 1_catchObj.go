package try

import (
	"encoding/json"
	"errors"
	"fmt"
)

type catchObj struct {
	logic func()
}

func (obj *catchObj) Catch(cb func(info *PanicInfo)) (result *PanicInfo) {
	if obj.logic == nil {
		return
	}
	defer func() {
		errInfo := recover()
		if errInfo != nil {
			if cb != nil {
				result = &PanicInfo{
					OriginObj: errInfo,
					ErrMsg:    obj.panicInfoToString(errInfo),
					Err:       obj.panicInfoToErr(errInfo),
					Stack:     readStack(),
				}
				cb(result)
			}
			return
		}
	}()
	obj.logic()
	return
}

func (obj *catchObj) DoNothing() *PanicInfo {
	return obj.Catch(nil)
}

func (obj *catchObj) Log() *PanicInfo {
	return obj.Catch(func(info *PanicInfo) {
		fmt.Println("PanicInfo", info.OriginObj)
		fmt.Println("ErrMsg", info.ErrMsg)
		fmt.Println("Stack", info.Stack)
	})
}

func (obj *catchObj) panicInfoToString(panicInfo interface{}) string {
	str, ok := panicInfo.(string)
	if ok {
		return str
	}
	content, err := json.Marshal(panicInfo)
	if err != nil {
		return fmt.Sprintf("%v", panicInfo)
	}
	return string(content)
}

func (obj *catchObj) panicInfoToErr(panicInfo interface{}) error {
	err, ok := panicInfo.(error)
	if ok {
		return err
	}
	return errors.New(obj.panicInfoToString(panicInfo))
}
