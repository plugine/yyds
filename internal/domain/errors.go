package domain

import "errors"

// 音乐解析和处理相关的错误
var (
	// ErrInvalidTempo 无效的节拍速度
	ErrInvalidTempo = errors.New("invalid tempo value")
	
	// ErrInvalidNote 无效的音符
	ErrInvalidNote = errors.New("invalid note symbol")
	
	// ErrInvalidPattern 无效的音乐模式
	ErrInvalidPattern = errors.New("invalid music pattern")
	
	// ErrEmptyMusic 空的音乐内容
	ErrEmptyMusic = errors.New("empty music content")
	
	// ErrSyntaxError 语法错误
	ErrSyntaxError = errors.New("syntax error in music file")
	
	// ErrUnsupportedFormat 不支持的格式
	ErrUnsupportedFormat = errors.New("unsupported music format")
)