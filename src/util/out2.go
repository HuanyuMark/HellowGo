package util

func GetDouble() float64 {
	testIfprotected()
	return float64(GetInt())
}
