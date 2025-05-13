package internal

type State int

const (
	List State = iota
	Detail
)

var state = List
