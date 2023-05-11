package stack

import (
	"fmt"
)

type Stack[T any] struct {
	Push  func(T)
	Pop   func() T
	Peek  func() T
	Retrn func() []T
	Print func()
}

func NewStack[T any]() stack[T] {
	var s stack[T]
	var st []T

	s.Push = func(i T) {
		st = append(st, i)
	}

	s.Pop = func() T {
		var val T

		slen := len(st)
		if slen < 1 {
			return val
		}

		val = st[slen-1]
		st = st[:slen-1]

		return val
	}

	s.Print = func() {
		fmt.Println(st)
	}

	s.Peek = func() T {
		if len(st) > 0 {
			return st[len(st)-1]
		}

		var t T
		return t
	}

	s.Retrn = func() []T {
		return st
	}

	return s
}
