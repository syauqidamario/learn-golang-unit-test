package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

//Assertion is to check for the exact value
//Testify is a Go-Lang unit testing library for assertion
func TestHelloWorldAssertion(t *testing.T){
	result := HelloWorld("Syauqi")
	assert.Equal(t, "Hi Syauqi", result)
	fmt.Println("Executed")
}

//assert would call the Fail() function is the checking fails.
//require would call the FailNow() function to not continue the unit test
func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Syauqi")
	require.Equal(t, "Hi Syauqi", result)
	fmt.Println("Executed")
}