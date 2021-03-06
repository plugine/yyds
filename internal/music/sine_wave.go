package music

import (
	"io"
	"math"
	"time"
)

const (
	sampleRate      = 44100
	channelNum      = 2
	bitDepthInBytes = 2
)

type SineWave struct {
	freq   float64
	length int64
	pos    int64

	remaining []byte
}

func timeBytesNum(duration time.Duration) int64 {
	l := int64(channelNum) * int64(bitDepthInBytes) * int64(sampleRate) * int64(duration) / int64(time.Second)
	return l
}

func NewSineWave(freq float64, duration time.Duration) *SineWave {
	l := timeBytesNum(duration)
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) == 0 {
		s.remaining = make([]byte, s.length)
		num := (bitDepthInBytes) * (channelNum)

		length := float64(sampleRate) / s.freq
		harmonics1Length := float64(sampleRate) / (s.freq * 2)
		harmonics2Length := float64(sampleRate) / (s.freq * 3)
		maxIndex := int(s.length) / num
		p := s.pos / int64(num)
		for i := 0; i < maxIndex; i++ {
			max := 32767.0 - (32767.0 * (float64(i) / float64(maxIndex)))
			fundamental := int16(math.Sin(2*math.Pi*float64(i)/float64(length)) * 0.6 * max)
			harmonics1 := int16(math.Sin(2*math.Pi*float64(i)/float64(harmonics1Length)) * 0.6 * max)
			harmonics2 := int16(math.Sin(2*math.Pi*float64(i)/float64(harmonics2Length)) * 0.6 * max)
			b := fundamental/2 + harmonics1/3 + harmonics2/4
			// b = fundamental
			for ch := 0; ch < channelNum; ch++ {
				s.remaining[num*i+2*ch] = byte(b)
				s.remaining[num*i+1+2*ch] = byte(b >> 8)
			}
			p++
		}
	}

	if s.pos >= s.length {
		return 0, io.EOF
	}
	n := copy(buf, s.remaining[s.pos:])
	s.pos += int64(n)
	return n, nil
}
