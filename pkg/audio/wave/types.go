package wave

import (
	"io"
	"time"
)

// Generator 波形生成器接口
type Generator interface {
	io.Reader
	// Duration 获取波形持续时间
	Duration() time.Duration
	// Frequency 获取波形频率
	Frequency() float64
}

// WaveType 波形类型
type WaveType int

const (
	// Sine 正弦波
	Sine WaveType = iota
	// Square 方波
	Square
	// Triangle 三角波
	Triangle
	// Sawtooth 锯齿波
	Sawtooth
)

// Config 波形生成配置
type Config struct {
	Frequency    float64       // 频率 (Hz)
	Duration     time.Duration // 持续时间
	SampleRate   int          // 采样率
	Amplitude    float64      // 振幅 (0.0-1.0)
	WaveType     WaveType     // 波形类型
}

// DefaultConfig 返回默认波形配置
func DefaultConfig(frequency float64, duration time.Duration) *Config {
	return &Config{
		Frequency:  frequency,
		Duration:   duration,
		SampleRate: 44100,
		Amplitude:  0.5,
		WaveType:   Sine,
	}
}