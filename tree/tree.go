package tree

type TraversalStrategy int

const (
	DepthFirst TraversalStrategy = iota
	BreadthFirst
)
