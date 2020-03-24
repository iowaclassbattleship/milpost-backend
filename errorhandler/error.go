package errorhandler

// ErrorHandler checks, if an error occured
func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
