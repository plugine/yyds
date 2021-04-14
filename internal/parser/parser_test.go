package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMusic = `
set tempo = 90

pattern(repeat = 4) {
    symphony:   e4 ++ d4 c4 b4 ++ c3 d3 a2 __ __ __ __ __ __ __
    background: f2 ++ ++ ++ g2 ++ ++ ++ a2 ++ ++ ++ ++ ++ ++ ++
}
`

func TestParseMusic(t *testing.T) {
	m := Parse(testMusic)
	fmt.Println(m)
	for k, v := range m.Patterns[0].TempoMap {
		fmt.Printf("%d ", k)
		for _, i := range v {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
	assert.Equal(t, 2, m.Patterns[0].TempoMap[6][0].TempoCount)
	assert.Equal(t, 120, m.Tempo)
	assert.Equal(t, 8, m.Patterns[0].AfterDelay)
}
