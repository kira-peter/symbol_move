# 符动世界 (SymbolMove)

字符符号在动，创造世界。

## 项目简介

SymbolMove 是一个用 Go 语言编写的终端艺术项目，旨在通过代码让字符"动起来"，创造各种酷炫的视觉效果。

## 特性

### 矩阵字符雨效果 🌧️

经典的"黑客帝国"风格字符雨效果，支持：

- ✨ 多种字符集（数字、字母、日文片假名、混合）
- 🎨 颜色渐变效果（亮白→亮绿→绿色→暗绿）
- ⚡ 可调节的下落速度（慢速/中速/快速）
- 📊 可调节的密度（稀疏/中等/密集）
- 🖥️ 自动适配终端尺寸
- 🎯 平滑的动画效果（可配置帧率）

## 快速开始

### 构建

```bash
go build -o matrix-rain.exe ./cmd/matrix-rain
```

### 运行

```bash
# 使用默认配置运行
./matrix-rain.exe

# 快速且密集的字符雨
./matrix-rain.exe -speed fast -density dense

# 慢速片假名字符雨
./matrix-rain.exe -charset katakana -speed slow

# 查看所有选项
./matrix-rain.exe -help
```

### 命令行选项

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

### 控制

- `ESC`, `Ctrl+C` 或 `q` - 退出程序

## 技术栈

- **Go** - 主要编程语言
- **tcell** - 跨平台终端控制库
- **标准库** - 尽量使用 Go 标准库

## 项目结构

```
symbol_move/
├── cmd/
│   └── matrix-rain/        # 矩阵字符雨可执行程序
│       └── main.go
├── pkg/
│   └── effects/
│       └── matrix-rain/    # 字符雨效果核心实现
│           └── rain.go
├── openspec/               # OpenSpec 规格和变更管理
│   ├── project.md          # 项目上下文
│   ├── AGENTS.md           # AI 助手指南
│   └── changes/            # 变更提案
└── go.mod
```

## 开发

本项目采用 OpenSpec 规范进行开发：

- 所有功能开发都从变更提案开始
- 规格驱动开发，确保需求明确
- 详见 `openspec/AGENTS.md`

## 许可证

MIT

## 路线图

- [x] 矩阵字符雨效果
- [ ] 更多终端特效...
- [ ] 字符动画系统...
- [ ] ASCII 艺术生成器...

---

**符动世界，让字符动起来！** ⚡
