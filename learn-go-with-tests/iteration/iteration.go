package iteration

const repeatCount = 5

func Repeat(str string) string {
	var repeat string
	for i := 0; i < repeatCount; i++ {
		repeat += "a"
	}
	return repeat
}