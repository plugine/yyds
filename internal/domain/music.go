package domain

import (
	"time"
	"yyds/pkg/music"
)

// Music 音乐作品的领域模型
type Music struct {
	Tempo    int                // BPM节拍
	Patterns []*music.Pattern   // 音乐模式列表
}

// NewMusic 创建新的音乐作品
func NewMusic() *Music {
	return &Music{
		Tempo:    120, // 默认120 BPM
		Patterns: make([]*music.Pattern, 0),
	}
}

// AddPattern 添加音乐模式
func (m *Music) AddPattern(pattern *music.Pattern) {
	m.Patterns = append(m.Patterns, pattern)
}

// SetTempo 设置节拍
func (m *Music) SetTempo(tempo int) {
	m.Tempo = tempo
}

// GetTempo16Duration 计算16分音符的持续时间
func (m *Music) GetTempo16Duration() time.Duration {
	// BPM是四分音符每分钟的数量
	// 16分音符 = 四分音符 / 4
	// 每分钟 = 60秒
	quarterNoteDuration := time.Minute / time.Duration(m.Tempo)
	return quarterNoteDuration / 4
}

// GetTotalDuration 计算音乐总时长
func (m *Music) GetTotalDuration() time.Duration {
	var totalDuration time.Duration
	tempo16Duration := m.GetTempo16Duration()
	
	for _, pattern := range m.Patterns {
		patternDuration := m.calculatePatternDuration(pattern, tempo16Duration)
		totalDuration += patternDuration * time.Duration(pattern.Repeat)
		totalDuration += tempo16Duration * time.Duration(pattern.AfterDelay)
	}
	
	return totalDuration
}

// calculatePatternDuration 计算单个模式的持续时间
func (m *Music) calculatePatternDuration(pattern *music.Pattern, tempo16Duration time.Duration) time.Duration {
	maxTime := pattern.GetMaxTime()
	return tempo16Duration * time.Duration(maxTime+1) // +1因为时间是从0开始的
}