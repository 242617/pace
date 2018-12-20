package utils

import (
	"math/rand"
	"time"
)

func NewGenerator(flags ...int) StringGenerator {

	charMap := []int{}
	src := []int{}
	for i := 0; i < 123; i++ {
		src = append(src, i)
	}
	smaz := src[97:123]
	lgaz := src[65:91]
	dgts := src[48:58]

	//       - 16  8  4  2  1  0
	// smaz	 1  0  1  0  1  0  1
	// lgaz  1  1  0  0  1  1  0
	// dgts  1  1  1  1  0  0  0

	if len(flags) != 0 {
		switch flags[0] {
		case 0:
			charMap = append(charMap, smaz...)
		case 1:
			charMap = append(charMap, lgaz...)
		case 2:
			charMap = append(charMap, smaz...)
			charMap = append(charMap, lgaz...)
		case 4:
			charMap = append(charMap, dgts...)
		case 8:
			charMap = append(charMap, dgts...)
			charMap = append(charMap, smaz...)
		case 16:
			charMap = append(charMap, dgts...)
			charMap = append(charMap, lgaz...)
		default:
			charMap = append(charMap, smaz...)
			charMap = append(charMap, lgaz...)
			charMap = append(charMap, dgts...)
		}
	} else {
		charMap = append(charMap, smaz...)
		charMap = append(charMap, dgts...)
	}

	return StringGenerator{charMap, int64(len(charMap)), 0}
}

type StringGenerator struct {
	charMap []int
	length  int64
	val     int64
}

func (g *StringGenerator) Generate(l int) (str string) {
	for i := 0; i < l; i++ {
		str += g.randSymbol()
	}
	return
}

func (g *StringGenerator) randSymbol() string {
	rand.Seed(time.Now().UnixNano() + g.nextValue())
	return string(g.charMap[rand.Int63n(g.length)])
}

func (g *StringGenerator) nextValue() int64 {
	(*g).val++
	if g.val > 9999 {
		(*g).val = 0
	}
	return g.val
}
