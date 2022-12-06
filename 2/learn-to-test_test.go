package learnToTest_test

import (
	"encoding/json"
	learnToTest "golang-learner/2"
	"testing"
)

func TestMyStuff_SetAddressWithoutPointer(t *testing.T) {
	// Write the test implementation to test the case where we are accessing the object directly and not a reference to the object:
	// TODO: Create an instance of MyStuff

	// TODO: Create an instance of Address fill in some fields

	// TODO: Invoke the function SetAddressWithoutPointer

	// TODO: Assert that the value returned from the accessor GetAddress() is nil
}

func TestMyStuff_SetAddressWithPointer(t *testing.T) {
	// Write the test implementation to test the case where we are accessing the reference to the object and not directly
	// TODO: Create an instance of MyStuff

	// TODO: Create an instance of Address fill in some fields

	// TODO: Invoke the function SetAddressWithPointer

	// TODO: Assert that the value returned from the accessor GetAddress() is not nil

	// TODO: Assert that the value returned from the accessor GetAddress() has the values you assigned at instantiation
}

func TestMyStuff(t *testing.T) {
	// Write the test implementation for the below JSON serialization cases (some code is already filled in)
	t.Run("Deserialization of the object omits the address field as it is private", func(t *testing.T) {
		// TODO: invoke the function getTestData() and assign it to a variable

		// TODO: Assert that the value returned from the accessor GetAddress() is nil

	})
	t.Run("Serialization/Deserialization of the object omits the Id field as it is marked to be skipped", func(t *testing.T) {
		//BEGIN Serialization
		// TODO: Create an instance of MyStuff and initialize the Id property to a non-nil value

		// TODO: use the json package to Marshal the MyStuff instance into a byte array and error

		// TODO: assert that the serialization was successful by asserting the error is nil

		// TODO: convert the byte array into a string and assert it does not have the "Id" property nor does it have the value you assigned it occurring anywhere

		// BEGIN Deserialization
		// TODO: invoke the function getTestData() and assign it to a variable

		// TODO: Assert that the Id field is nil

	})
	t.Run("Serialization of the object omits the Name field when it is the empty value", func(t *testing.T) {
		//BEGIN Serialization
		data := learnToTest.MyStuff{}
		_ = data // This is just here to allow the file to compile by mocking that data is used
		// TODO: use the json package to Marshal the MyStuff instance into a byte array and error

		// TODO: assert that the serialization was successful by asserting the error is nil

		// TODO: convert the byte array into a string and assert it does not have the "name" property occurring anywhere
	})

	t.Run("Serialization of the object omits the Name field when it is the empty value", func(t *testing.T) {
		//BEGIN Serialization
		data := learnToTest.MyStuff{Age: 11}
		_ = data // This is just here to allow the file to compile by mocking that data is used
		// TODO: use the json package to Marshal the MyStuff instance into a byte array and error

		// TODO: assert that the serialization was successful by asserting the error is nil

		// TODO: convert the byte array into a string and assert it has the property "someAge" with value 11
	})
}

func getTestData() learnToTest.MyStuff {
	serialiedJSON := "{\"Id\": \"secret id\", \"name\":\"my coolest name\", \"address\": {\"StreetNumber\": 12, \"StreetName\": \"street\", \"City\": \"minneapolis\", \"State\": \"MN\"}, \"someAge\": 3}"
	var data learnToTest.MyStuff
	_ = json.Unmarshal([]byte(serialiedJSON), &data) //This is the way to ignore one or more of the returned parameters from a function
	return data
}