package try

type catchObj struct {
	logic func()
}

func (obj *catchObj) Catch(cb func(er interface{})) {
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

func Try(logic func()) *catchObj {
	return &catchObj{
		logic: logic,
	}
}
