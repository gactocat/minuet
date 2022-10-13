package pkg

type Block int

const (
	__ Block = iota
	NM
	SP
)

func (b Block) String() string {
	switch b {
	case __:
		return "＋"
	case NM:
		return "⬜"
	case SP:
		return "🔶"
	default:
		return "Error"
	}
}
