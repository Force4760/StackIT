package stack

import "math"

func (s *Stack) Plus() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	s.Push(a + b)
	return nil
}

func (s *Stack) Minus() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	s.Push(a - b)
	return nil
}

func (s *Stack) Times() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	s.Push(a * b)
	return nil
}

func (s *Stack) Div() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	if b == 0 {
		return ErrDivideBy0
	}
	s.Push(a / b)
	return nil
}

func (s *Stack) Modulus() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	if b == 0 {
		return ErrDivideBy0
	}
	modValue := float64(int(a) % int(b))
	s.Push(modValue)
	return nil
}

func (s *Stack) Power() error {
	if s.length < 2 {
		return ErrNotEnoughtValuesOnTheStack
	}
	a, b := s.Pop2NoTest()
	powerValue := math.Pow(a, b)
	s.Push(powerValue)
	return nil
}
