package oto

import (
	"fmt"
	"yyds/pkg/audio"
	
	"github.com/hajimehoshi/oto"
)

// Engine oto音频引擎实现
type Engine struct {
	context *oto.Context
	config  *audio.Config
}

// NewEngine 创建新的oto音频引擎
func NewEngine(config *audio.Config) *Engine {
	if config == nil {
		config = audio.DefaultConfig()
	}
	
	return &Engine{
		config: config,
	}
}

// Initialize 实现 audio.Engine 接口
func (e *Engine) Initialize() error {
	ctx, err := oto.NewContext(e.config.SampleRate, e.config.Channels, e.config.BitDepth/8, e.config.BufferSize)
	if err != nil {
		return err
	}
	
	e.context = ctx
	return nil
}

// Play 实现 audio.Engine 接口
func (e *Engine) Play(data []byte) error {
	if e.context == nil {
		return fmt.Errorf("audio engine not initialized")
	}
	
	player := e.context.NewPlayer()
	defer player.Close()
	
	_, err := player.Write(data)
	return err
}

// Close 实现 audio.Engine 接口
func (e *Engine) Close() error {
	if e.context != nil {
		return e.context.Close()
	}
	return nil
}