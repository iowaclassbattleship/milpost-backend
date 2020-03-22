package api

// ErrorHandler does
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
