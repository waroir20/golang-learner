package queue

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//-------------------THIS IS THE INTERFACE YOU WILL IMPLEMENT----------------------------
//---------------------------------------------------------------------------------------

// Queue is the representation of the simple FIFO (First In First Out) data structure
type Queue interface {
	// Contains - checks if the element is in the Queue
	Contains(string) bool
	// Push - adds an element to the back of the Queue
	Push(string)
	// Pop - removes the first element from the Queue and returns its value, return error if Queue empty
	Pop() (string, error)
	// IsEmpty - returns true if there are no elements in the Queue, false otherwise.
	IsEmpty() bool
	// Size - returns the number of elements in the Queue
	Size() int
}
