# SymbolMove 特效快速参考

## 运行主程序

```bash
# 编译
go build -o symbol-move.exe ./cmd/symbol-move

# 运行（显示特效选择菜单）
./symbol-move.exe
```

## 操作说明

### 主菜单
- **↑/↓ 或 j/k**: 移动选择
- **Enter**: 运行选中的特效
- **q 或 Ctrl+C**: 退出程序

### 特效运行时
- **ESC**: 退出特效，返回主菜单

## 13个特效说明

### 1. 音频可视化 (audio-visualizer)
频谱柱状图效果，模拟音频波形展示
- 标签: 音频、可视化、频谱、动画

### 2. 大字时钟 (big-clock)
ASCII Art 大字体显示当前时间 (HH:MM:SS)
- 标签: 时钟、实用、ASCII艺术、时间
- 自动更新: 每秒

### 3. 数字瀑布 (digital-waterfall)
数字 0-9 从上到下快速流动，绿色主题
- 标签: 科技、动画、绿色、经典
- 风格: 类似黑客帝国

### 4. 火焰燃烧 (fire-effect)
从底部向上燃烧的火焰效果
- 标签: 火焰、热、动画、自然
- 颜色: 红黄渐变

### 5. 生命游戏 (game-of-life)
Conway's Game of Life 细胞自动机
- 标签: 算法、自动机、经典、模拟
- 算法: 经典生命游戏规则

### 6. 矩阵字符雨 (matrix-rain)
经典黑客帝国风格字符雨效果
- 标签: classic、matrix、animation、green
- 风格: 绿色科技

### 7. 迷宫生成 (maze-generator)
实时展示 DFS 算法生成迷宫过程
- 标签: 算法、迷宫、生成、路径
- 算法: 深度优先搜索

### 8. 粒子爆炸 (particle-burst)
从中心点向外爆炸的粒子效果
- 标签: 粒子、爆炸、动画、物理
- 物理: 重力模拟、衰减效果

### 9. Plasma 等离子 (plasma)
彩色等离子云效果，正弦函数生成图案
- 标签: 数学、彩色、等离子、科幻
- 算法: 正弦波叠加

### 10. 彩虹波浪 (rainbow-wave)
彩虹色的水平波浪从左到右滚动
- 标签: 彩色、波浪、动画、美观
- 颜色: 7种彩虹色

### 11. 雪花飘落 (snowfall)
浪漫的雪花从天而降
- 标签: 自然、冬日、浪漫、动画
- 效果: 多层飘落

### 12. 星空闪烁 (starry-sky)
模拟夜空中星星闪烁的效果
- 标签: 自然、动画、简单、放松
- 效果: 正弦波闪烁

### 13. 波浪文字 (wave-text)
显示 "SymbolMove" 文字以正弦波形式波动
- 标签: 文字、动画、彩色、流畅
- 效果: 彩虹渐变

## 技术规格

### 性能
- 帧率: 30 FPS（大部分特效）
- 自适应: 自动适配终端大小
- 优化: 高效渲染，低 CPU 占用

### 架构
- 接口: 统一 Effect 接口
- 注册: 自动注册机制
- 扩展: 易于添加新特效

## 开发指南

### 添加新特效

1. 创建目录结构：
```bash
mkdir -p pkg/effects/[effect-name]
mkdir -p openspec/changes/add-batch-effects/specs/[effect-name]
```

2. 创建核心文件：
- `pkg/effects/[effect-name]/[core].go` - 核心逻辑
- `pkg/effects/[effect-name]/effect.go` - Effect 接口实现
- `pkg/effects/[effect-name]/register.go` - 自动注册
- `openspec/changes/add-batch-effects/specs/[effect-name]/spec.md` - 规格文档

3. 在 `cmd/symbol-move/main.go` 中添加导入：
```go
_ "github.com/symbolmove/symbol_move/pkg/effects/[effect-name]"
```

### 必须实现的接口

```go
type Effect interface {
    Metadata() Metadata
    Init(screen tcell.Screen) error
    Run(quit <-chan struct{}) error
    Cleanup() error
}
```

## 项目结构

```
symbol_move/
├── cmd/
│   └── symbol-move/
│       └── main.go                    # 主程序
├── pkg/
│   ├── effects/
│   │   ├── effect.go                  # Effect 接口定义
│   │   ├── registry.go                # 注册表
│   │   ├── audio-visualizer/          # 音频可视化
│   │   ├── big-clock/                 # 大字时钟
│   │   ├── digital-waterfall/         # 数字瀑布
│   │   ├── fire-effect/               # 火焰燃烧
│   │   ├── game-of-life/              # 生命游戏
│   │   ├── matrix-rain/               # 矩阵字符雨
│   │   ├── maze-generator/            # 迷宫生成
│   │   ├── particle-burst/            # 粒子爆炸
│   │   ├── plasma/                    # Plasma 等离子
│   │   ├── rainbow-wave/              # 彩虹波浪
│   │   ├── snowfall/                  # 雪花飘落
│   │   ├── starry-sky/                # 星空闪烁
│   │   └── wave-text/                 # 波浪文字
│   └── ui/
│       └── selector/                  # 特效选择器 UI
└── openspec/
    └── changes/
        └── add-batch-effects/
            └── specs/                 # 所有特效规格文档
```

## 常见问题

### Q: 如何退出特效？
A: 按 ESC 键

### Q: 如何退出程序？
A: 在主菜单按 q 或 Ctrl+C

### Q: 特效运行卡顿怎么办？
A: 减小终端窗口大小，或关闭其他占用 CPU 的程序

### Q: 如何自定义特效参数？
A: 目前使用默认配置，未来版本将支持命令行参数配置

## 许可证

本项目遵循开源许可证（待定）
