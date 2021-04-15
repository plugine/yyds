# YYDS - 目标是变成专业的文本音乐编辑器

现在还极端不专业，作者乐理基础不过关，还在持续学习和优化中

![otto suwen](pics/otto.jpg)

使用了 [oto](https://github.com/hajimehoshi/oto) 作为音频播放库，所以仓库命名为`YYDS-永远滴神`

## 示例
```text
// Young And Beautiful

// 设定节拍
set tempo = 120

// 顺序播放一段音乐(`pattern`)，一段音乐由多个音轨组成，音轨由标识和一系列音符记号组成
pattern(repeat = 1) {
    // 小写记号为 major 调(如 `a4`)，大写记号为 minor 调(如 `E4`)
    main:  b4{2} ++ a4{2} ++ d5 b4{2} ++ a4{2} ++ d5 b4{2} ++ a4{2} ++ ++ g4 ++ ++ ++
    // `++` 表示前一个音符延音一个全音，`__` 表示占位，一个全音时间内不发声
    back:  f3 ++ ++ ++ g3 ++ ++ ++ a3 ++ ++ ++ ++ ++ ++ ++
    // 花括号表示非全音，如 `a3{2}` 表示`a3`为二分音符，只发声半个节拍；'*'表示相同音重复几次，如 `a3*2` 等同于 `a3 a3`
    human: __*8 __ __ __ a3{2} b3{2} c4{2} e4{2} b3{2} c4{2} ++*2
}

pattern {
    main: a4{2} a4 b4{2} b4 ++ a4 b4{2}   b4 ++ a4 d5       b4 ++ a4 ++ g4
    back:       __ __    f3 ++*3          g3 ++*2 ++{2}           a3 ++*6
    human: __*8 __ __ __{2} a3{2} b3{2} c4{2} e4{2} b3{2} c4{2} ++*2
}
```

以上就是全部的语法了，更多示例可以查看 example 文件夹。

## 运行
```shell
git clone https://github.com/plugine/yyds.git
cd yyds
# linux 需要安装libasound2-dev库，apt install libasound2-dev
# 如果需要修改语法，需要安装 antlr4，并在变更 Music.g4 后重新生成编译文件
go run cmd/yyds/main.go example/shutdown.yyds
```

## TODOLIST

- [x] 语法解析
- [x] 音乐播放
- [x] 正弦波添加泛音，模拟电子琴音
- [ ] 探索更多乐器