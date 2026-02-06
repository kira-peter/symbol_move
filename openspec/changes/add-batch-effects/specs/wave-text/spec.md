# 波浪文字特效

## 概述

波浪文字特效显示一段文字，文字以正弦波形式上下波动，并配有彩色渐变效果，创造出动态流畅的视觉效果。

## 特效 ID

`wave-text`

## 显示名称

波浪文字

## 描述

显示文字以正弦波形式上下波动，配有彩色渐变效果

## 核心特性

### 1. 波浪动画

- **正弦波运动**: 每个字符沿正弦波路径上下移动
- **相位差**: 相邻字符有固定相位差，形成波浪效果
- **平滑动画**: 使用浮点数计算，确保平滑移动

### 2. 颜色渐变

- **彩虹渐变**: 字符颜色从左到右呈现彩虹渐变
- **动态变化**: 颜色随时间循环变化
- **丰富色彩**: 使用 RGB 颜色系统

### 3. 可配置文本

- **默认文本**: "SymbolMove"
- **可自定义**: 支持配置任意文本
- **居中显示**: 文本在屏幕中央显示

## 技术实现

### 数据结构

```go
type WaveText struct {
    screen    tcell.Screen
    text      string
    phase     float64  // 波浪相位
    colorPhase float64  // 颜色相位
    width     int
    height    int
}
```

### 核心算法

```go
// 计算字符 Y 坐标
y = centerY + amplitude * sin(phase + x * phaseShift)

// 计算字符颜色
hue = (colorPhase + x * colorShift) % 360
```

## 配置参数

```go
type Config struct {
    Text       string  // 显示文本
    Amplitude  float64 // 波浪振幅
    WaveSpeed  float64 // 波浪速度
    ColorSpeed float64 // 颜色变化速度
    FPS        int     // 帧率
}
```

## 特效元数据

- **ID**: wave-text
- **名称**: 波浪文字
- **标签**: ["文字", "动画", "彩色", "流畅"]
