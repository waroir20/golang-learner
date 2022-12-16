package stack_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-learner/1/stack"
	"testing"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//-----------------THESE TESTS ARE WHAT YOU ARE BUILDING AGAINST-------------------------
//---------------------------------------------------------------------------------------

func TestIsEmpty(t *testing.T) {
	// Since push does not have a return we must test it some way else
	t.Run("returns true for empty stack", func(t *testing.T) {
		workerB := stack.New()
		assert.True(t, workerB.IsEmpty())
	})
	t.Run("returns false after single push", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		assert.False(t, workerB.IsEmpty())
	})
	t.Run("returns false after double push", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		workerB.Push("monkey")
		assert.False(t, workerB.IsEmpty())
	})
}

func TestPop(t *testing.T) {
	t.Run("returns error when stack empty", func(t *testing.T) {
		workerB := stack.New()
		_, err := workerB.Pop()
		assert.NotNil(t, err)
	})
	t.Run("returns top element of stack", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		element, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, "banana", element)
	})
	t.Run("returns last element pushed onto stack", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("0")
		workerB.Push("nuts")
		workerB.Push("bolts")
		workerB.Push("banana")
		element, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, "banana", element)
	})
	t.Run("returns error after all elements removed", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		_, err := workerB.Pop()
		assert.Nil(t, err)
		_, err = workerB.Pop()
		assert.NotNil(t, err)
	})
}

func TestSize(t *testing.T) {
	t.Run("returns 0 for empty stack", func(t *testing.T) {
		workerB := stack.New()
		assert.Equal(t, 0, workerB.Size())
	})
	t.Run("returns 1 after single push", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		assert.Equal(t, 1, workerB.Size())
	})
	t.Run("returns 2 after double push", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		workerB.Push("banana")
		assert.Equal(t, 2, workerB.Size(), "duplicate elements should not matter")
	})
	t.Run("returns 0 after all elements removed", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		_, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 0, workerB.Size())
	})
}

func TestContains(t *testing.T) {
	t.Run("returns true when element in stack", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		assert.True(t, workerB.Contains("banana"))
	})
	t.Run("returns true when element is not top of stack", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("0")
		workerB.Push("banana")
		workerB.Push("nuts")
		workerB.Push("bolts")
		assert.True(t, workerB.Contains("banana"))
	})
	t.Run("returns false when element not in stack", func(t *testing.T) {
		workerB := stack.New()
		workerB.Push("banana")
		assert.False(t, workerB.Contains("jabber"))
	})
	t.Run("returns false when no elements in stack", func(t *testing.T) {
		workerB := stack.New()
		assert.False(t, workerB.Contains("jabber"))
	})
}

func TestPush(t *testing.T) {
	// Within reason your data structure should be able to handle mostly unlimited number of elements
	workerB := stack.New()
	for i := 0; i < 10_000; i++ {
		workerB.Push(fmt.Sprintf("%d", i))
		assert.Equal(t, i, workerB.Size()-1)
	}
	assert.True(t, workerB.Contains("0"), "can access bottom of stack")
	assert.False(t, workerB.Contains("XYZ"))
}
