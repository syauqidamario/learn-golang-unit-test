package helper

import (
	"fmt"
	"testing"
)

//Error() would automatically the Fail() function to consider the unit test fail.
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Syauqi")
	if result != "Hi Syauqi"{
		//failed unit test
		// panic("Result is not Hi Syauqi")
		// error using t.Error
		t.Error("It's supposed to be Hi Syauqi")
	}
	fmt.Println("Executed")
}

//Fatal() would do an error log by calling the FailNow() function to stop the unit test execution
func TestHelloWorldFatal(t *testing.T){
	result := HelloWorld("Syauqi")
	if result != "Hi Syauqi"{
		t.Fatal("It's supposed to be Hi Syauqi")
	}
	fmt.Println("Executed")
}