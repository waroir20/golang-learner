package queue

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//-------------------THIS IS THE INTERFACE YOU WILL IMPLEMENT----------------------------
//---------------------------------------------------------------------------------------

// Queue is the representation of the simple FIFO (First In First Out) data structure
type Queue[E comparable] interface {
	// Contains - checks if the element is in the Queue
	Contains(E) bool
	// Push - adds an element to the back of the Queue
	Push(E)
	// Pop - removes the first element from the Queue and returns its value, return error if Queue empty
	Pop() (E, error)
	// IsEmpty - returns true if there are no elements in the Queue, false otherwise.
	IsEmpty() bool
	// Size - returns the number of elements in the Queue
	Size() int
}
