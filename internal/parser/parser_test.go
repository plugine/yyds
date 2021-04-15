package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMusic(t *testing.T) {
	var testMusic = `
		set tempo = 90

		pattern(repeat = 4) {
			symphony:   e4 ++ d4 c4 b4 ++ c3 d3 a2 __ __ __ __ __ __ __
			background: f2 ++ ++ ++ g2 ++ ++ ++ a2 ++ ++ ++ ++ ++ ++ ++
		}`
	m := Parse(testMusic)
	for k, v := range m.Patterns[0].TempoMap {
		fmt.Printf("%d ", k)
		for _, i := range v {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
	assert.Equal(t, 2, m.Patterns[0].TempoMap[5][0].Tempo16Count)
	assert.Equal(t, 90, m.Tempo)
	assert.Equal(t, 4, m.Patterns[0].Repeat)
}

func TestNoteRepeatParse(t *testing.T) {
	testMusic := `
		set tempo = 90
		pattern {
			notes: a1{2} ++{2}*3
		}`
	m := Parse(testMusic)
	for k, v := range m.Patterns[0].TempoMap {
		fmt.Printf("%d ", k)
		for _, i := range v {
			fmt.Printf("%v", i)
		}
		fmt.Printf("\n")
	}
}
