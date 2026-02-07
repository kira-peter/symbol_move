# 符动世界 (SymbolMove)

字符符号在动，创造世界。

## 项目简介

SymbolMove 是一个用 Go 语言编写的终端艺术项目，旨在通过代码让字符"动起来"，创造各种酷炫的视觉效果。

## 快速开始

### 构建

```bash
# 构建主程序（推荐）
go build -o symbol-move.exe ./cmd/symbol-move

# 或者构建独立的特效程序
go build -o matrix-rain.exe ./cmd/matrix-rain
```

### 运行

```bash
# 启动主界面（推荐方式）
./symbol-move.exe

# 使用上下键选择特效，回车确认运行
# ESC 返回主界面，q 退出程序
```

**独立运行特效**：
```bash
# 直接运行矩阵字符雨
./matrix-rain.exe

# 使用参数自定义
./matrix-rain.exe -speed fast -density dense -charset katakana
```

## 特性

### 🎯 交互式特效选择器

- **可视化菜单** - 清晰的界面浏览所有特效
- **键盘导航** - 上下键选择，回车确认,ESC 返回
- **多语言界面** - 支持中英文切换（Ctrl+Space），自动保存语言偏好
- **特效预览** - 显示每个特效的描述信息
- **即时切换** - 快速在不同特效之间切换

### ✨ 13 种精彩特效

#### 自然类特效
- **🌧️ 矩阵字符雨** - 经典黑客帝国风格，多种字符集和颜色渐变
- **⭐ 星空闪烁** - 夜空中星星随机闪烁，支持多种颜色主题
- **❄️ 雪花飘落** - 冬日雪景，三层深度效果，自然摆动动画

#### 数据流特效
- **🔢 数字瀑布** - 数字 0-9 快速流动，绿色主题

#### 文字动画
- **🌊 波浪文字** - 文字以正弦波形式波动，彩虹渐变
- **🌈 彩虹波浪** - 七彩波浪从左到右滚动

#### 粒子系统
- **💥 粒子爆炸** - 多点爆炸效果，物理模拟，重力和衰减

#### 实用工具
- **🕐 大字时钟** - ASCII Art 大字体显示当前时间

#### 高级算法
- **🔥 火焰燃烧** - 热量传播算法，红黄渐变火焰
- **🧬 生命游戏** - Conway's Game of Life 细胞自动机
- **🌀 迷宫生成** - DFS 算法实时生成迷宫动画
- **🎨 Plasma 等离子** - 彩色等离子云效果
- **🎵 音频可视化** - 频谱柱状图，动态波形展示

## 使用指南

### 主界面操作

```
符动世界 (SymbolMove)
字符符号在动，创造世界
─────────────────────────────────────
  1. ► 矩阵字符雨
  2.   星空闪烁
  3.   雪花飘落
  4.   数字瀑布
  ...更多特效...
─────────────────────────────────────
描述：经典黑客帝国风格字符雨效果
─────────────────────────────────────
↑↓: 选择  |  Enter: 确认  |  q: 退出
```

**键盘控制**：
- `↑` / `k` - 向上选择
- `↓` / `j` - 向下选择
- `←` / `h` - 向左列移动
- `→` / `l` - 向右列移动
- `Enter` - 运行选中的特效
- `ESC` - 从特效返回主界面
- `T` 或 `t` - 切换界面语言（中文/English）
- `q` / `Ctrl+C` - 退出程序
- `1-9` / `0` - 数字快捷键直接选择

### 矩阵字符雨选项

仅在使用独立程序 `matrix-rain.exe` 时可用：

```
-speed string
      下落速度: slow, medium, fast (默认 medium)
-density string
      字符雨密度: sparse, medium, dense (默认 medium)
-charset string
      字符集: digits, letters, katakana, mixed (默认 mixed)
-fps int
      帧率 (默认 30)
-help
      显示帮助信息
```

示例：
```bash
./matrix-rain.exe -speed fast -density dense
./matrix-rain.exe -charset katakana -speed slow
```

## 技术栈

- **Go** - 主要编程语言
- **tcell** - 跨平台终端控制库
- **OpenSpec** - 规格驱动开发流程

## 项目结构

```
symbol_move/
├── cmd/
│   ├── symbol-move/         # 主程序入口（特效选择器）
│   │   └── main.go
│   └── matrix-rain/         # 矩阵字符雨独立程序
│       └── main.go
├── pkg/
│   ├── effects/             # 特效系统
│   │   ├── effect.go        # 特效接口定义
│   │   ├── registry.go      # 特效注册表
│   │   ├── matrix-rain/     # 矩阵字符雨
│   │   ├── starry-sky/      # 星空闪烁
│   │   ├── snowfall/        # 雪花飘落
│   │   ├── digital-waterfall/ # 数字瀑布
│   │   ├── wave-text/       # 波浪文字
│   │   ├── rainbow-wave/    # 彩虹波浪
│   │   ├── particle-burst/  # 粒子爆炸
│   │   ├── big-clock/       # 大字时钟
│   │   ├── fire-effect/     # 火焰燃烧
│   │   ├── game-of-life/    # 生命游戏
│   │   ├── maze-generator/  # 迷宫生成
│   │   ├── plasma/          # Plasma 等离子
│   │   └── audio-visualizer/ # 音频可视化
│   └── ui/
│       └── selector/        # 选择器 UI 组件
│           └── selector.go
├── openspec/                # OpenSpec 规格和变更管理
│   ├── project.md           # 项目上下文
│   ├── AGENTS.md            # AI 助手指南
│   └── changes/             # 变更提案
└── go.mod
```

## 架构设计

### 特效系统

所有特效实现统一的 `Effect` 接口：

```go
type Effect interface {
    Metadata() Metadata        // 元数据
    Init(screen) error         // 初始化
    Run(quit <-chan) error     // 运行
    Cleanup() error            // 清理
}
```

**特点**：
- 插件化架构 - 新特效只需实现接口并注册
- 生命周期管理 - Init → Run → Cleanup
- 统一的错误处理和资源清理
- 支持热插拔（无需修改主程序代码）

### 添加新特效

1. 在 `pkg/effects/` 下创建新目录
2. 实现 `Effect` 接口
3. 在 `init()` 中注册：`effects.Register(NewEffect)`
4. 自动出现在主界面中

## 开发

本项目采用 OpenSpec 规范进行开发：

- 所有功能开发都从变更提案开始
- 规格驱动开发，确保需求明确
- 详见 `openspec/AGENTS.md`

## 许可证

MIT

## 路线图

- [x] 矩阵字符雨效果
- [x] 特效选择器主界面
- [x] 星空闪烁
- [x] 雪花飘落
- [x] 数字瀑布
- [x] 波浪文字动画
- [x] 彩虹波浪
- [x] 粒子爆炸特效
- [x] 大字时钟
- [x] 火焰燃烧效果
- [x] 生命游戏（Conway's Game of Life）
- [x] 迷宫生成动画
- [x] Plasma 等离子效果
- [x] 音频可视化
- [ ] ASCII 艺术生成器
- [ ] 终端分形图案
- [ ] 更多创意特效...

---

**符动世界，让字符动起来！** ⚡
