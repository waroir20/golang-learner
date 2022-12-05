package queue_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-learner/1/queue"
	"testing"
)

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//-----------------THESE TESTS ARE WHAT YOU ARE BUILDING AGAINST-------------------------
//---------------------------------------------------------------------------------------

func TestIsEmpty(t *testing.T) {
	// Since Push does not have a return we must test it some way else
	t.Run("returns true for empty queue", func(t *testing.T) {
		workerB := queue.New()
		assert.True(t, workerB.IsEmpty())
	})
	t.Run("returns false after single push", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		assert.False(t, workerB.IsEmpty())
	})
	t.Run("returns false after double push", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		workerB.Push("monkey")
		assert.False(t, workerB.IsEmpty())
	})
}

func TestPop(t *testing.T) {
	t.Run("returns error when queue empty", func(t *testing.T) {
		workerB := queue.New()
		_, err := workerB.Pop()
		assert.NotNil(t, err)
	})
	t.Run("returns first element from queue", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		element, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, "banana", element)
	})
	t.Run("returns first element pushed into queue", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		workerB.Push("bolts")
		workerB.Push("nuts")
		workerB.Push("0")
		element, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, "banana", element)
	})
	t.Run("returns error after all elements removed", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		_, err := workerB.Pop()
		assert.Nil(t, err)
		_, err = workerB.Pop()
		assert.NotNil(t, err)
	})
}

func TestSize(t *testing.T) {
	t.Run("returns 0 for empty queue", func(t *testing.T) {
		workerB := queue.New()
		assert.Equal(t, 0, workerB.Size())
	})
	t.Run("returns 1 after single push", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		assert.Equal(t, 1, workerB.Size())
	})
	t.Run("returns 2 after double push", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		workerB.Push("banana")
		assert.Equal(t, 2, workerB.Size(), "duplicate elements should not matter")
	})
	t.Run("returns 0 after all elements removed", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		_, err := workerB.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 0, workerB.Size())
	})
}

func TestContains(t *testing.T) {
	t.Run("returns true when element in queue", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		assert.True(t, workerB.Contains("banana"))
	})
	t.Run("returns true when element is not first in queue", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("0")
		workerB.Push("banana")
		workerB.Push("nuts")
		workerB.Push("bolts")
		assert.True(t, workerB.Contains("banana"))
	})
	t.Run("returns false when element not in queue", func(t *testing.T) {
		workerB := queue.New()
		workerB.Push("banana")
		assert.False(t, workerB.Contains("jabber"))
	})
	t.Run("returns false when no elements in queue", func(t *testing.T) {
		workerB := queue.New()
		assert.False(t, workerB.Contains("jabber"))
	})
}

func TestPush(t *testing.T) {
	// Within reason your data structure should be able to handle mostly unlimited number of elements
	workerB := queue.New()
	for i := 0; i < 10_000; i++ {
		workerB.Push(fmt.Sprintf("%d", i))
		assert.Equal(t, i, workerB.Size())
	}
	assert.True(t, workerB.Contains("10000"), "can access back of queue")
	assert.False(t, workerB.Contains("XYZ"))
}
