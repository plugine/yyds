package config

import (
	"yyds/pkg/audio"
)

// Config 应用程序配置
type Config struct {
	Audio  *AudioConfig  `json:"audio"`
	Parser *ParserConfig `json:"parser"`
	Debug  bool          `json:"debug"`
}

// AudioConfig 音频相关配置
type AudioConfig struct {
	Engine     string `json:"engine"`      // 音频引擎类型 (oto, etc.)
	SampleRate int    `json:"sample_rate"` // 采样率
	Channels   int    `json:"channels"`    // 声道数
	BitDepth   int    `json:"bit_depth"`   // 位深度
	BufferSize int    `json:"buffer_size"` // 缓冲区大小
}

// ParserConfig 解析器相关配置
type ParserConfig struct {
	StrictMode bool `json:"strict_mode"` // 严格模式
}

// DefaultConfig 返回默认配置
func DefaultConfig() *Config {
	return &Config{
		Audio: &AudioConfig{
			Engine:     "oto",
			SampleRate: 44100,
			Channels:   2,
			BitDepth:   16,
			BufferSize: 1024,
		},
		Parser: &ParserConfig{
			StrictMode: false,
		},
		Debug: false,
	}
}

// ToAudioConfig 转换为音频引擎配置
func (c *AudioConfig) ToAudioConfig() *audio.Config {
	return &audio.Config{
		SampleRate: c.SampleRate,
		Channels:   c.Channels,
		BitDepth:   c.BitDepth,
		BufferSize: c.BufferSize,
	}
}