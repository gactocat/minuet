package pkg

const (
	UP Dir = iota
	LEFT
	DOWN
	RIGHT
)

var (
	AllDir = []Dir{UP, LEFT, DOWN, RIGHT}
)

type Dir int
