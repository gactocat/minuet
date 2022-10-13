package pkg

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExist = errors.New("already exist")
)

type Field [][]Block

func CreateField(w int, h int) Field {
	ret := make(Field, h)
	for y := 0; y < h; y++ {
		ret[y] = make([]Block, w)
	}
	return ret
}

func (f Field) Put(p Point, c Card) (Field, error) {
	ret := f.Copy()
	for y, bb := range c.Blocks {
		for x, b := range bb {
			if b == __ {
				// 空ブロックは配置しない
				continue
			}
			if ret[p.Y+y][p.X+x] != __ {
				// 配置先が空でなければエラー
				return nil, ErrAlreadyExist
			}
			ret[p.Y+y][p.X+x] = b
		}
	}
	return ret, nil
}

func (f Field) GetSidePoints() []Point {
	pp := map[Point]bool{}
	for y, bb := range f {
		for x, b := range bb {
			if b == __ {
				continue
			}

			for offsetY := -1; offsetY <= +1; offsetY++ {
				for offsetX := -1; offsetX <= +1; offsetX++ {
					if f[y+offsetY][x+offsetX] == __ {
						pp[Point{X: x + offsetX, Y: y + offsetY}] = true
					}
				}
			}
		}
	}

	ret := make([]Point, 0, len(pp))
	for p := range pp {
		ret = append(ret, p)
	}
	return ret
}

func (f Field) GetSpecialPoint() int {
	p := 0
	for y, bb := range f {
	Block:
		for x, b := range bb {
			if b == SP {
				for offsetY := -1; offsetY <= +1; offsetY++ {
					for offsetX := -1; offsetX <= +1; offsetX++ {
						if f[y+offsetY][x+offsetX] == __ {
							continue Block
						}
					}
				}
				p++
			}
		}
	}
	return p
}

func (f Field) Width() int {
	max := 0
	for _, bb := range f {
		l := len(bb)
		if l > max {
			max = l
		}
	}
	return max
}

func (f Field) Height() int {
	return len(f)
}

func (f Field) Copy() Field {
	ret := make(Field, len(f))
	for y, bb := range f {
		ret[y] = make([]Block, len(bb))
		for x, b := range bb {
			ret[y][x] = b
		}
	}
	return ret
}

func (f Field) String() string {
	s := ""
	for x := 0; x < f.Width(); x++ {
		s += fmt.Sprintf("%2d", x)
	}
	s += "\n"
	for y, bb := range f {
		for _, b := range bb {
			s += b.String()
		}
		s += fmt.Sprintf("%2d\n", y)
	}
	return s
}
