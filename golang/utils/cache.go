package utils

func Cache(f func() interface{}) func() interface{} {
	var m interface{}

	return func() interface{} {
		if m == nil {
			m = f()
		}
		return m
	}
}
