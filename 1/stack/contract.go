package stack

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//-------------------THIS IS THE INTERFACE YOU WILL IMPLEMENT----------------------------
//---------------------------------------------------------------------------------------

// Stack is the representation of the simple LIFO (Last In First Out) data structure
type Stack interface {
	// Contains - checks if the element is in the Stack
	Contains(string) bool
	// IsEmpty - returns true if there are no elements in the Stack, false otherwise.
	IsEmpty() bool
	// Pop - removes the top element from the Stack and returns its value, return error if Stack empty
	Pop() (string, error)
	// Push - adds an element to the top of the Stack
	Push(string)
	// Size - returns the number of elements in the Stack
	Size() int
}
