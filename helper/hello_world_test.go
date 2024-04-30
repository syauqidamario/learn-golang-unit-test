package helper

import (
	"fmt"
	"runtime"
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

func TestSkip(t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("Unit test can't run on Windows")
	}
	result := HelloWorld("Syauqi")
	require.Equal(t, "Hi Syauqi", result)
}

//testing.M parameter is to manage an unit test's execution
//if a TestMain function exists, it would automatically execute the function when a unit test runs in a package
func TestMain(m *testing.M){
	fmt.Println("Before Unit Test")
	m.Run() //this will execute ALL unit test
	fmt.Println("After Unit Test")
}

//Go-Lang supports the creation of unit test functions inside a unit test functions using the Run() function
func TestSubTest(t *testing.T){
	t.Run("Syauqi", func(t *testing.T){
		result := HelloWorld("Syauqi")
		require.Equal(t, "Hi Syauqi", result)
	})
	t.Run("Damario Djohan", func(t *testing.T){
		result := HelloWorld("Damario Djohan")
		require.Equal(t, "Hi Damario Djohan", result)
	})
}