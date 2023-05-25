package try

func Try(logic func()) *catchObj {
	return &catchObj{
		logic: logic,
	}
}

type PanicInfo struct {
	OriginObj interface{}
	ErrMsg    string
	Err       error
	Stack     string
}
