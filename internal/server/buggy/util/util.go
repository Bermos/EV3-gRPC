package util

func GetDefaultString(val, d string) string {
	if val == "" {
		return d
	}

	return val
}

func GetDefaultNumber[t int | float32](val, d t) t {
	if val == 0 {
		return d
	}

	return val
}
