package learnToTest_test

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	learnToTest "golang-learner/2"
	"testing"
)

// Learning Goal: understand that member functions belonging to an object
// at the value level have access to the data but cannot modify it
// -> Essentially the member function is passing the value of the instance of the struct to the function in the background
func TestMyStuff_SetAddressWithoutPointer(t *testing.T) {
	// Write the test implementation to test the case where we are accessing the object directly and not a reference to the object:
	// TODO: Create an instance of MyStuff

	// TODO: Create an instance of Address fill in some fields

	// TODO: Invoke the function SetAddressWithoutPointer

	// TODO: Assert that the value returned from the accessor GetAddress() is nil
}

// Learning Goal: understand that member functions belonging to an object
// at the reference level have access to the data and can modify it
// -> Essentially the member function is passing the reference to the instance of the struct to the function in the background
func TestMyStuff_SetAddressWithPointer(t *testing.T) {
	// Write the test implementation to test the case where we are accessing the reference to the object and not directly
	// TODO: Create an instance of MyStuff

	// TODO: Create an instance of Address fill in some fields

	// TODO: Invoke the function SetAddressWithPointer

	// TODO: Assert that the value returned from the accessor GetAddress() is not nil

	// TODO: Assert that the value returned from the accessor GetAddress() has the values you assigned at instantiation
}

// Learning goal: understand a common struct tag "json" and its role in deserialization/serialization
func TestMyStuff(t *testing.T) {
	// Write the test implementation for the below JSON serialization cases (some code is already filled in)
	t.Run("Deserialization of the object omits the address field as it is private", func(t *testing.T) {
		// TODO: invoke the function getTestMyStuff() and assign it to a variable

		// TODO: Assert that the value returned from the accessor GetAddress() is nil

	})
	t.Run("Serialization/Deserialization of the object omits the Id field as it is marked to be skipped", func(t *testing.T) {
		//BEGIN Serialization
		// TODO: Create an instance of MyStuff and initialize the Id property to a non-nil value

		// TODO: use the json package to Marshal the MyStuff instance into a byte array and error

		// TODO: assert that the serialization was successful by asserting the error is nil

		// TODO: convert the byte array into a string and assert it does not have the "Id" property nor does it have the value you assigned it occurring anywhere

		// BEGIN Deserialization
		// TODO: invoke the function getTestMyStuff() and assign it to a variable

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

	t.Run("Serialization of the object sets the Age field as the json property 'someAge'", func(t *testing.T) {
		//BEGIN Serialization
		data := learnToTest.MyStuff{Age: 11}
		_ = data // This is just here to allow the file to compile by mocking that data is used
		// TODO: use the json package to Marshal the MyStuff instance into a byte array and error

		// TODO: assert that the serialization was successful by asserting the error is nil

		// TODO: convert the byte array into a string and assert it has the property "someAge" with value 11
	})
}

func getTestMyStuff() learnToTest.MyStuff {
	serializedJSON := "{\"Id\": \"secret id\", \"name\":\"my coolest name\", \"address\": {\"StreetNumber\": 12, \"StreetName\": \"street\", \"City\": \"minneapolis\", \"State\": \"MN\"}, \"someAge\": 3}"
	var data learnToTest.MyStuff
	_ = json.Unmarshal([]byte(serializedJSON), &data) //This is the way to ignore one or more of the returned parameters from a function
	return data
}

// Learning Goal: understand golang struct validation:
// using library - https://github.com/go-playground/validator
// NOTE: also test your ability to fill in missing steps
func TestAddress(t *testing.T) {
	modelValidator := validator.New()
	t.Run("Validation passes for value returned by getTestAddress()", func(t *testing.T) {
		err := modelValidator.Struct(getTestAddress())
		assert.Nil(t, err)
	})
	t.Run("Validation of the object when Id is marked uuid", func(t *testing.T) {
		// TODO: use the modelValidator to show that when the Id value is not a UUID that validation returns an error
	})
	t.Run("Validation of the object when StreetNumber is marked greater than 0", func(t *testing.T) {
		// TODO: use the modelValidator to show that when the StreetNumber value is negative that validation returns an error
	})
	t.Run("Validation of the object when StreetName is marked required", func(t *testing.T) {
		// TODO: use the modelValidator to show that when the StreetName is empty that validation returns an error
	})
	t.Run("Validation of the object when State is marked required if Country is USA", func(t *testing.T) {
		// TODO: use the modelValidator to show that when the Country is USA and state is empty that validation returns an error
	})
	t.Run("Validation of the object when PostalCode is marked as min length of 2", func(t *testing.T) {
		// TODO: use the modelValidator to show that when the PostalCode less than 2 characters that validation returns an error
	})
}

func getTestAddress() learnToTest.Address {
	return learnToTest.Address{
		Id:           "61e3ef7d-6af9-4105-a84f-d671b916ab60",
		StreetNumber: 234,
		StreetName:   "Lopporit avenue",
		City:         "Livingway",
		State:        "",
		Country:      "UK",
		PostalCode:   "12",
	}
}
