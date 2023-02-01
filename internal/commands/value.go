package commands

import "time"

type Value struct {
	Next Command
}

func (v Value) GetType() CommandType {
	return ValueType
}

func (v Value) GetNext() Command {
	return v.Next
}

type Money struct {
	Value
	Val float64
}

type Dates struct {
	Value
	From time.Time
	To time.Time
}