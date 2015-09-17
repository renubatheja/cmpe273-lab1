package main

//Go Test file for unit test cases written for MySleep.go 

import ("testing";"time")

func TestSleepFiveSeconds(t *testing.T) {
	oldNow := time.Now().Second()
	Sleep(5)
	currNow := time.Now().Second()
	diffInSeconds := currNow - oldNow
	if diffInSeconds != 5 {
		t.Error("Expected 5, but got ", diffInSeconds)
	}
}

