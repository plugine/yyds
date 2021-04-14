package music

import (
	"encoding/binary"
	"time"

	"github.com/hajimehoshi/oto"
)

const (
	ContinueSymbol = Symbol("++")
	SilentSymbol   = Symbol("__")
)

var (
	mixedBytes []byte
)

type Symbol string

func (s Symbol) Freq() float64 {
	return freqMap[string(s)]
}

func mix(bytes [][]byte) {
	tempBytes := make([]byte, 2)
	for i := 0; i < len(mixedBytes)/2; i++ {
		it := int16(0)
		for _, b := range bytes {
			tempBytes[0], tempBytes[1] = b[i*2+1], b[i*2]
			it = int16(binary.BigEndian.Uint16(tempBytes))/int16(2) + it
		}
		mixedBytes[i*2] = byte(it)
		mixedBytes[i*2+1] = byte(it >> 8)
	}
}

type Note struct {
	Symbol           Symbol
	TempoCount       int // 延音几个节拍
	NoteSign         int // 是什么音符，默认为全音符
	CurrentNoteIndex int // 当前播放到第几个音符

	wave *SineWave
}

func (n *Note) buffer(tempoDuration time.Duration) []byte {
	if n.wave == nil {
		n.wave = NewSineWave(n.Symbol.Freq(), tempoDuration*time.Duration(n.TempoCount))
	}
	buf := make([]byte, timeBytesNum(tempoDuration))
	_, err := n.wave.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}

type Pattern struct {
	Repeat        int
	AfterDelay    int
	MaxTempoCount int             // 小节最大节拍数
	TempoMap      map[int][]*Note // 每个节拍对应的音符
}

func (p *Pattern) Play(tempoDuration time.Duration, player *oto.Player) {
	// repeat
	for i := 0; i <= p.Repeat; i++ {
		// 清空已读状态
		for _, notelist := range p.TempoMap {
			for _, note := range notelist {
				if note.wave != nil {
					note.wave.pos = 0
				}
			}
		}
		for i := 0; i < p.MaxTempoCount; i++ {
			var bytes [][]byte
			if currentNoteLines, exists := p.TempoMap[i]; exists {
				for _, note := range currentNoteLines {
					bytes = append(bytes, note.buffer(tempoDuration))
				}
			}
			mix(bytes)
			player.Write(mixedBytes)
		}
		// after delay
		time.Sleep(tempoDuration * time.Duration(p.AfterDelay))
	}
}

type Music struct {
	Tempo    int
	Patterns []Pattern
}

func (m Music) Play() {
	tempoDuration := time.Duration(60.0/float64(m.Tempo)*1000) * time.Millisecond
	mixedBytes = make([]byte, timeBytesNum(tempoDuration))
	c, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, 4096)
	player := c.NewPlayer()
	if err != nil {
		panic(err)
	}

	for _, p := range m.Patterns {
		p.Play(tempoDuration, player)
	}
	player.Close()
}
