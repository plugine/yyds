package music

import (
	"time"
)

// Symbol 音符符号类型
type Symbol string

// 特殊音符符号常量
const (
	ContinueSymbol Symbol = "++" // 延音符号
	SilentSymbol   Symbol = "__" // 静音符号
)

// Note 音符结构
type Note struct {
	Symbol       Symbol        // 音符符号
	Tempo16Count int          // 持续多少个16分音符 (1-16)
	Frequency    float64      // 音符频率
}

// NewNote 创建新音符
func NewNote(symbol Symbol, tempo16Count int) *Note {
	note := &Note{
		Symbol:       symbol,
		Tempo16Count: tempo16Count,
	}
	
	// 设置频率
	if freq, exists := GetFrequency(string(symbol)); exists {
		note.Frequency = freq
	}
	
	return note
}

// IsSpecial 检查是否为特殊音符（延音或静音）
func (n *Note) IsSpecial() bool {
	return n.Symbol == ContinueSymbol || n.Symbol == SilentSymbol
}

// IsContinue 检查是否为延音符
func (n *Note) IsContinue() bool {
	return n.Symbol == ContinueSymbol
}

// IsSilent 检查是否为静音符
func (n *Note) IsSilent() bool {
	return n.Symbol == SilentSymbol
}

// Duration 计算音符的实际持续时间
func (n *Note) Duration(tempo16Duration time.Duration) time.Duration {
	return tempo16Duration * time.Duration(n.Tempo16Count)
}

// GetFrequency 获取音符频率
func (s Symbol) Freq() float64 {
	if freq, exists := GetFrequency(string(s)); exists {
		return freq
	}
	return 0.0
}