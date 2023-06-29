package recursive

func FibD(n int) (res int) {
	if n == 1 {
		return 1
	} else if n > 1 {
		res = n * FibD(n-1)
	}
	return
}
