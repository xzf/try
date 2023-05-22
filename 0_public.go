package try

func Try(logic func()) *catchObj {
	return &catchObj{
		logic: logic,
	}
}
