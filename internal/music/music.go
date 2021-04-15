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
	Symbol       Symbol
	Tempo16Count int // 声音持续多少个 16 分音符，最大为16，最小为1

	wave *SineWave
}

func (n *Note) buffer(tempo16Duration time.Duration) []byte {
	if n.wave == nil {
		n.wave = NewSineWave(n.Symbol.Freq(), tempo16Duration*time.Duration(n.Tempo16Count))
	}
	buf := make([]byte, timeBytesNum(tempo16Duration))
	_, err := n.wave.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}

type Pattern struct {
	Repeat          int
	AfterDelay      int             // 播放完当前小节后，延迟多少 16分音符再播放下一段
	MaxTempo16Count int             // 小节最大16分音符数
	TempoMap        map[int][]*Note // 每个16分节拍对应的音符
}

func (p *Pattern) Play(tempo16Duration time.Duration, player *oto.Player) {
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
		for i := 0; i < p.MaxTempo16Count; i++ {
			var bytes [][]byte
			if currentNoteLines, exists := p.TempoMap[i]; exists {
				for _, note := range currentNoteLines {
					bytes = append(bytes, note.buffer(tempo16Duration))
				}
			}
			mix(bytes)
			player.Write(mixedBytes)
		}
		// after delay
		time.Sleep(tempo16Duration * time.Duration(p.AfterDelay))
	}
}

type Music struct {
	Tempo    int
	Patterns []Pattern
}

func (m Music) Play() {
	// 以16分音符为最小播放单位
	tempo16Duration := time.Duration(60.0/float64(m.Tempo)*1000) * time.Millisecond / 16
	mixedBytes = make([]byte, timeBytesNum(tempo16Duration))
	c, err := oto.NewContext(sampleRate, channelNum, bitDepthInBytes, 4096)
	player := c.NewPlayer()
	if err != nil {
		panic(err)
	}

	for _, p := range m.Patterns {
		p.Play(tempo16Duration, player)
	}
	player.Close()
}
