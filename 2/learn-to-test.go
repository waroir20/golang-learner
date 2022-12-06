package learnToTest

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

type MyStuff struct {
	Id      *string  `json:"-"`
	Name    string   `json:"name,omitempty"`
	address *Address `json:"address"` //THIS is correct it looks like an error because it doesn't actually work
	Age     int      `json:"someAge"`
}

type Address struct {
	StreetNumber int
	StreetName   string `validate:"required,"` //TODO continue here
	City         string
	State        string
	Country      string
}

// SetAddressWithoutPointer - Notice how this does not update the address field
func (m MyStuff) SetAddressWithoutPointer(address Address) {
	m.address = &address
}

// SetAddressWithPointer - Notice how this does update the address field
func (m *MyStuff) SetAddressWithPointer(address Address) {
	m.address = &address
}

// GetAddress - Simple accessor for the package private field 'address'
func (m MyStuff) GetAddress() *Address {
	return m.address
}
