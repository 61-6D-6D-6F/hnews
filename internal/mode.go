package internal

type Mode int

const (
	List Mode = iota
	Details
	Comments
)

var mode = List
