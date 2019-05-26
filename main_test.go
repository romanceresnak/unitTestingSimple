package main

import "testing"

func TestHelloGo(t *testing.T) {
		//expected value
		expected := "Welcome to go language"
		//current value from main program
		actual := HelloGo()
		//compare this 2 values
		if actual != expected{
			t.Error("Test failed",expected,actual)
	}
}
