package stack

import (
	"fmt"
)

type Stack[T any] struct {
	push  func(T)
	pop   func() T
	peek  func() T
	retrn func() []T
	print func()
}

func NewStack[T any]() stack[T] {
	var s stack[T]
	var st []T

	s.push = func(i T) {
		st = append(st, i)
	}

	s.pop = func() T {
		var val T

		slen := len(st)
		if slen < 1 {
			return val
		}

		val = st[slen-1]
		st = st[:slen-1]

		return val
	}

	s.print = func() {
		fmt.Println(st)
	}

	s.peek = func() T {
		if len(st) > 0 {
			return st[len(st)-1]
		}

		var t T
		return t
	}

	s.retrn = func() []T {
		return st
	}

	return s
}
