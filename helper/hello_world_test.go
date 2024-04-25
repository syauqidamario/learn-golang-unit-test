package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syauqi")
	if result != "Hi Syauqi"{
		//failed unit test
		panic("Result is not Hi Syauqi")
	}
}