package recursion

func Fibonacci(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
