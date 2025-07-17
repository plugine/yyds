# YYDS 项目重构方案

## 当前项目分析

这是一个音乐文本编辑器项目，主要功能包括：
- 解析音乐文本语法 (使用 ANTLR)
- 生成音频波形
- 播放音乐

## 当前项目结构问题

1. **缺乏清晰的包分层**：所有代码都在 `internal` 下，没有明确的业务逻辑分层
2. **包职责不够清晰**：parser 包中包含了语法文件生成的代码和业务逻辑混在一起
3. **缺少 API/接口层**：没有清晰的对外接口
4. **缺乏配置管理**：配置分散在代码中
5. **缺少 pkg 目录**：没有可重用的公共库
6. **测试文件组织不规范**：测试覆盖不完整

## 建议的新项目结构

```
yyds/
├── cmd/                    # 应用程序入口
│   └── yyds/
│       └── main.go
├── pkg/                    # 可重用的公共库
│   ├── audio/             # 音频处理相关
│   │   ├── engine.go      # 音频引擎接口
│   │   ├── oto/           # oto 音频引擎实现
│   │   │   └── engine.go
│   │   └── wave/          # 波形生成
│   │       ├── sine.go
│   │       └── types.go
│   ├── music/             # 音乐理论相关
│   │   ├── frequency.go   # 频率计算
│   │   ├── note.go        # 音符定义
│   │   ├── pattern.go     # 模式定义
│   │   └── tempo.go       # 节拍相关
│   └── player/            # 播放器相关
│       ├── player.go      # 播放器接口和实现
│       └── mixer.go       # 音频混合
├── internal/              # 内部包，不对外暴露
│   ├── app/               # 应用层
│   │   ├── config/        # 配置管理
│   │   │   └── config.go
│   │   └── service/       # 业务服务层
│   │       ├── composer.go    # 音乐合成服务
│   │       └── parser.go      # 解析服务
│   ├── domain/            # 领域模型
│   │   ├── music.go       # 音乐领域模型
│   │   └── errors.go      # 错误定义
│   └── infrastructure/    # 基础设施层
│       ├── antlr/         # ANTLR 生成的代码
│       │   ├── music_base_listener.go
│       │   ├── music_lexer.go
│       │   ├── music_listener.go
│       │   ├── music_parser.go
│       │   └── Music.g4
│       └── parser/        # 解析器实现
│           ├── listener.go
│           └── parser.go
├── examples/              # 示例文件
│   ├── basic/
│   │   └── simple.yyds
│   ├── advanced/
│   │   ├── doubletiger.yyds
│   │   ├── shutdown.yyds
│   │   └── youngandbeautiful.yyds
│   └── README.md
├── test/                  # 集成测试
│   ├── testdata/         # 测试数据
│   └── integration/      # 集成测试代码
├── docs/                  # 文档
│   ├── API.md            # API 文档
│   ├── GRAMMAR.md        # 语法说明
│   └── DEVELOPMENT.md    # 开发指南
├── scripts/              # 脚本文件
│   ├── build.sh          # 构建脚本
│   └── generate.sh       # 代码生成脚本
├── .gitignore
├── go.mod
├── go.sum
├── Makefile              # 构建配置
└── README.md
```

## 重构优势

### 1. 清晰的分层架构
- **cmd/**: 应用程序入口，保持简洁
- **pkg/**: 可重用的公共库，可以被其他项目引用
- **internal/**: 内部实现，不对外暴露

### 2. 职责分离
- **app/**: 应用层，处理配置和业务流程
- **domain/**: 领域模型，包含核心业务逻辑
- **infrastructure/**: 基础设施，包含外部依赖

### 3. 包设计原则
- 每个包有明确的职责
- 包之间依赖关系清晰
- 遵循依赖倒置原则

### 4. 可测试性
- 接口分离，便于 Mock 测试
- 测试文件组织清晰
- 支持单元测试和集成测试

## 具体重构步骤

### 第一阶段：创建新的目录结构
1. 创建 `pkg/` 相关目录
2. 创建 `internal/` 新的分层目录
3. 移动现有文件到合适位置

### 第二阶段：重构代码
1. 定义清晰的接口
2. 分离音频引擎和音乐逻辑
3. 重构解析器代码
4. 添加配置管理

### 第三阶段：完善项目
1. 添加构建脚本
2. 完善文档
3. 添加更多测试
4. 优化包的导入关系

## 接口设计

### 音频引擎接口
```go
package audio

type Engine interface {
    Initialize() error
    Play(data []byte) error
    Close() error
}
```

### 音乐播放器接口
```go
package player

type Player interface {
    LoadMusic(music *domain.Music) error
    Play() error
    Stop() error
    Pause() error
}
```

### 解析器接口
```go
package service

type Parser interface {
    Parse(content string) (*domain.Music, error)
    Validate(content string) error
}
```

## 依赖管理

新的项目结构将依赖关系组织为：
```
cmd -> internal/app -> internal/domain
                   -> pkg/
internal/infrastructure -> pkg/
```

这样确保了：
- `pkg/` 包不依赖 `internal/` 包
- 核心业务逻辑在 `domain` 中
- 外部依赖隔离在 `infrastructure` 中