## 1. 语言管理模块 (pkg/i18n)
- [x] 1.1 创建 `pkg/i18n/language.go` - 定义语言类型和管理器
- [x] 1.2 创建 `pkg/i18n/translations.go` - 定义翻译键和翻译映射
- [x] 1.3 创建 `pkg/i18n/config.go` - 实现配置文件读写(保存到 ~/.symbolmove/config.json)
- [x] 1.4 实现语言切换和获取当前语言的方法

## 2. 特效元数据国际化 (pkg/effects)
- [x] 2.1 扩展 `Metadata` 结构,添加 `NameEN` 和 `DescriptionEN` 字段
- [x] 2.2 更新所有特效的 `effect.go` 文件,添加英文名称和描述:
  - [x] matrix-rain (矩阵字符雨)
  - [x] starry-sky (星空效果)
  - [x] snowfall (雪花飘落)
  - [x] digital-waterfall (数字瀑布)
  - [x] wave-text (波浪文字)
  - [x] rainbow-wave (彩虹波浪)
  - [x] particle-burst (粒子爆发)
  - [x] big-clock (大字时钟)
  - [x] fire-effect (火焰效果)
  - [x] game-of-life (生命游戏)
  - [x] maze-generator (迷宫生成)
  - [x] plasma (等离子体)
  - [x] audio-visualizer (音频可视化)
  - [x] heartbeat (心跳特效)
  - [x] dna-helix (DNA 螺旋)
  - [x] ocean-wave (海洋波浪)
  - [x] fireworks (烟花效果)
  - [x] water-ripple (水波涟漪)
  - [x] matrix-tunnel (矩阵隧道)
  - [x] typewriter-code (打字机代码)
  - [x] tetris-auto (俄罗斯方块 AI)
  - [x] snake-ai (贪吃蛇 AI)
  - [x] qrcode-gen (二维码生成器)

## 3. 选择器界面国际化 (pkg/ui/selector)
- [x] 3.1 修改 `Selector` 结构,添加语言管理器引用
- [x] 3.2 修改 `renderTitle()` - 支持翻译标题和副标题
- [x] 3.3 修改 `renderEffectList()` - 根据当前语言显示特效名称
- [x] 3.4 修改 `renderDescription()` - 根据当前语言显示特效描述
- [x] 3.5 修改 `renderHints()` - 支持翻译操作提示
- [x] 3.6 添加语言指示器显示(在标题区域显示 "中文/English")
- [x] 3.7 在 `HandleKey()` 中添加 `Ctrl+Space` 快捷键处理逻辑

## 4. 主程序集成 (cmd/symbol-move)
- [x] 4.1 在 `main()` 中初始化语言系统(加载配置)
- [x] 4.2 将语言管理器传递给选择器
- [x] 4.3 确保语言切换后保存配置

## 5. 测试和验证
- [x] 5.1 测试 `Ctrl+Space` 快捷键切换功能
- [x] 5.2 测试所有界面文本的翻译正确性
- [x] 5.3 测试语言偏好持久化(重启后保持)
- [x] 5.4 测试各个特效名称和描述的中英文显示
- [ ] 5.5 测试在不同终端尺寸下的显示效果

## 6. 文档更新
- [x] 6.1 更新 README.md 添加语言切换说明
- [x] 6.2 添加配置文件说明文档

