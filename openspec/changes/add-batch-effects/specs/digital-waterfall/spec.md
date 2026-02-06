# 数字瀑布特效

## 概述

数字瀑布特效是一个类似黑客帝国矩阵雨的视觉效果，但只使用数字字符(0-9)从上到下快速流动，营造出数字瀑布的视觉效果。

## 特效 ID

`digital-waterfall`

## 显示名称

数字瀑布

## 描述

数字 0-9 从上到下快速流动，绿色主题的数字瀑布效果

## 核心特性

### 1. 数字流

- **随机数字**: 只使用 0-9 十个数字字符
- **列流动**: 每列独立流动
- **速度变化**: 不同列有不同的流动速度
- **尾迹效果**: 数字后方有渐变的尾迹

### 2. 颜色系统

- **主色调**: 绿色主题（绿色科技感）
- **亮度渐变**:
  - 头部: 亮绿色（最新的数字）
  - 中部: 中绿色
  - 尾部: 暗绿色（逐渐消失）
- **随机高亮**: 偶尔出现白色数字作为闪光

### 3. 性能优化

- **帧率**: 30 FPS 流畅运行
- **自适应**: 自动适配终端大小
- **内存管理**: 固定长度的数字流，避免内存泄漏

## 技术实现

### 数据结构

```go
type Column struct {
    x         int     // 列的 x 坐标
    y         int     // 当前头部的 y 坐标
    speed     float64 // 流动速度（单位：格/帧）
    length    int     // 数字流长度
    digits    []rune  // 数字序列
    spawnDelay int    // 生成延迟（帧数）
}

type DigitalWaterfall struct {
    screen     tcell.Screen
    columns    []*Column
    width      int
    height     int
    rand       *rand.Rand
}
```

### 核心算法

1. **初始化**: 为每列创建随机速度和长度的数字流
2. **更新逻辑**:
   - 移动每列的 y 坐标
   - 生成新的随机数字
   - 移除超出屏幕的数字
3. **渲染逻辑**:
   - 根据位置计算亮度
   - 应用颜色渐变
   - 绘制数字字符

### 颜色配置

```go
// 绿色渐变色板
var greenShades = []tcell.Color{
    tcell.ColorLime,      // 最亮
    tcell.ColorGreen,     // 亮
    tcell.ColorDarkGreen, // 中
    tcell.ColorGray,      // 暗
}
```

## 用户交互

- **ESC**: 退出特效，返回主界面
- **终端调整**: 自动适应终端大小变化

## 配置参数

```go
type Config struct {
    MinSpeed   float64 // 最小流动速度 (默认: 0.5)
    MaxSpeed   float64 // 最大流动速度 (默认: 2.0)
    MinLength  int     // 最小流长度 (默认: 5)
    MaxLength  int     // 最大流长度 (默认: 15)
    FPS        int     // 帧率 (默认: 30)
}
```

## 特效元数据

- **ID**: digital-waterfall
- **名称**: 数字瀑布
- **描述**: 数字 0-9 从上到下快速流动，绿色主题的数字瀑布效果
- **标签**: ["科技", "动画", "绿色", "经典"]
- **版本**: 1.0.0

## 文件结构

```
pkg/effects/digital-waterfall/
├── waterfall.go    # 核心逻辑实现
├── effect.go       # Effect 接口实现
└── register.go     # 自动注册
```

## 测试要点

1. 验证数字流平滑流动
2. 验证颜色渐变效果
3. 验证 ESC 键正常退出
4. 验证终端大小调整后正常显示
5. 验证 30 FPS 流畅运行
