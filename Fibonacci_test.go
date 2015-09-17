package main

//Go Test file for unit test cases written for Fibonacci.go 

import ("testing")

func TestFibonacciZero(t *testing.T) {
	var num int64
	num = 0
	num = Fibonacci(num)
	if num != 0 {
		t.Error("Expected 0, but got ", num)
	}
}

func TestFibonacciOne(t *testing.T) {
	var num int64
	num = 1
	num = Fibonacci(num)
	if num != 1 {
		t.Error("Expected 1, but got ", num)
	}
}


func TestFibonacciHugeNumber(t *testing.T) {
	var num int64
	num = 40
	num = Fibonacci(num)
	if num != 102334155 {
		t.Error("Expected 102334155, but got ", num)
	}
}

func TestFibonacciMultipleNumbers(t *testing.T) {
	var fibTests = []struct {
  		n        int64 // input
  		expected int64 // expected result
	}{
	  {1, 1},
	  {2, 1},
	  {3, 2},
	  {4, 3},
	  {5, 5},
	  {6, 8},
	  {7, 13},
	}
	
	for _, tt := range fibTests {
    	actual := Fibonacci(tt.n)
	    if actual != tt.expected {
	      t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
	    }
	}	
}

func TestFibonacciNegative(t *testing.T) {
	var num int64
	num = -5
	num = Fibonacci(num)
	if num != -1 {
		t.Error("Expected -1, but got ", num)
	}
}
