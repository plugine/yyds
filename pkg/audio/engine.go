package audio

// Engine 定义音频引擎接口
type Engine interface {
	// Initialize 初始化音频引擎
	Initialize() error
	
	// Play 播放音频数据
	Play(data []byte) error
	
	// Close 关闭音频引擎
	Close() error
}

// Config 音频引擎配置
type Config struct {
	SampleRate   int // 采样率
	Channels     int // 声道数
	BitDepth     int // 位深度
	BufferSize   int // 缓冲区大小
}

// DefaultConfig 返回默认音频配置
func DefaultConfig() *Config {
	return &Config{
		SampleRate: 44100,
		Channels:   2,
		BitDepth:   16,
		BufferSize: 1024,
	}
}