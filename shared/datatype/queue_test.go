package datatype_test

import (
	"testing"

	"github.com/n0rmanc/fthelper/shared/datatype"
	"github.com/n0rmanc/fthelper/shared/xtests"
)

func TestNormalQueue(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("create normal queue").
		WithExpected(0).
		WithActual(datatype.NewQueue().Size()).
		MustEqual()

	assertion.NewName("create limit queue").
		WithExpected(2).
		WithActual(datatype.
			NewLimitQueue(2).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Size()).
		MustEqual()

	assertion.NewName("correct head on limit queue").
		WithExpected(3).
		WithActual(datatype.
			NewLimitQueue(3).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Head()).
		MustEqual()

	assertion.NewName("correct tail on limit queue").
		WithExpected(5).
		WithActual(datatype.
			NewLimitQueue(3).
			Enqueue(1).
			Enqueue(2).
			Enqueue(3).
			Enqueue(4).
			Enqueue(5).
			Tail()).
		MustEqual()

	var queue = datatype.NewLimitQueue(3).
		Enqueue(1).
		Enqueue(2).
		Enqueue(3)

	assertion.NewName("should able to get first queue").
		WithExpected(1).
		WithActual(queue.Get()).
		MustEqual()

	assertion.NewName("should able to remove queue").
		WithExpected(2).
		WithActual(queue.Size()).
		MustEqual()

	var queue1 = datatype.NewLimitQueue(0)

	assertion.NewName("should not able to get first queue").
		WithExpected(nil).
		WithActual(queue1.Get()).
		MustEqual()

	assertion.NewName("should return true if it is empty").
		WithExpected(true).
		WithActual(datatype.NewLimitQueue(0).Empty()).
		MustEqual()

	assertion.NewName("should return false if it is not empty").
		WithExpected(false).
		WithActual(datatype.NewLimitQueue(1).Enqueue(1).Empty()).
		MustEqual()

	var array = []interface{}{1, 2}

	assertion.NewName("should convert array to queue").
		WithExpected(2).
		WithActual(datatype.ToQueue(array).Size()).
		MustDeepEqual()
}
