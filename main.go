package main

import (
	"fmt"
	"time"

	"github.com/gactocat/minuet/pkg"
)

func main() {
    t := time.Now()
	result := map[int]map[int]pkg.Field{}

	for i, c1 := range pkg.Cards {
		f := pkg.CreateField(30, 30)
		f, _ = f.Put(pkg.Pos{X: 12, Y: 12}, c1)
		for j := i + 1; j < len(pkg.Cards); j++ {
			c2 := pkg.Cards[j]
			for _, fp := range f.GetSidePoints() {
				for _, d := range pkg.AllDir {
					tc := c2.Turn(d)
					for _, cp := range tc.GetFilledPoint() {
						p := pkg.Pos{X: fp.X - cp.X, Y: fp.Y - cp.Y}
						f, err := f.Put(p, tc)
						if err != nil {
							continue
						}
						if f.GetSpecialPoint() < 2 {
							continue
						}
						if _, ok := result[c1.No]; !ok {
							result[c1.No] = map[int]pkg.Field{}
						}
						result[c1.No][c2.No] = f
					}
				}
			}
		}
	}

	for c1, m := range result {
		for c2, f := range m {
			fmt.Printf("Card1:%v\n%v", c1, pkg.NoCardMap[c1])
			fmt.Printf("Card2:%v\n%v", c2, pkg.NoCardMap[c2])
			fmt.Printf("%v", f)
		}
	}
	fmt.Printf("Count: %v", len(result))
    fmt.Printf("Time: %vs\n", time.Since(t).Seconds())
}
