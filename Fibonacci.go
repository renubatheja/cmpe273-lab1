package main


func Fibonacci(x int64) int64 {
		if(x < 0) {
			return -1
		} else {
			if x == 0 {
				return 0
			} else if x == 1 {
				return 1
			} else {
				return Fibonacci(x-1) + Fibonacci(x-2)
			}
		}
}