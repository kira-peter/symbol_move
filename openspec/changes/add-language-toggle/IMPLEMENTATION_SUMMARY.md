# 语言切换功能实施总结

## 概述

成功为 SymbolMove 项目添加了中英文界面切换功能,用户可通过 `Ctrl+Space` 快捷键在中英文之间切换,语言偏好会自动保存到配置文件。

## 实施内容

### 1. 创建 i18n 国际化模块 (`pkg/i18n/`)

#### language.go
- 定义 `Language` 类型和常量 (`LanguageChinese`, `LanguageEnglish`)
- 实现线程安全的 `Manager` 单例模式
- 提供语言切换、查询等核心方法

#### translations.go
- 定义翻译键常量 (KeyTitle, KeySubtitle, KeyHints 等)
- 建立中英文翻译映射表
- 实现 `T()` 函数用于翻译查询
- 提供 `GetEffectName()` 和 `GetEffectDescription()` 辅助函数

#### config.go
- 实现配置文件读写功能
- 配置保存到 `~/.symbolmove/config.json`
- 提供优雅的错误处理和默认值回退

### 2. 扩展特效元数据 (`pkg/effects/effect.go`)

- 为 `Metadata` 结构添加 `NameEN` 和 `DescriptionEN` 字段
- 保持向后兼容,原有 `Name` 和 `Description` 字段用于中文
- 所有 23 个特效文件都已更新英文翻译:
  * matrix-rain, starry-sky, snowfall, digital-waterfall
  * wave-text, rainbow-wave, particle-burst, big-clock
  * fire-effect, game-of-life, maze-generator, plasma
  * audio-visualizer, heartbeat, dna-helix, ocean-wave
  * fireworks, water-ripple, matrix-tunnel, typewriter-code
  * tetris-auto, snake-ai, qrcode-gen

### 3. 选择器界面国际化 (`pkg/ui/selector/selector.go`)

- 导入 i18n 包
- 修改 `renderTitle()` 使用翻译后的标题和副标题
- 修改 `renderEffectList()` 根据语言显示特效名称
- 修改 `renderDescription()` 根据语言显示特效描述
- 修改 `renderHints()` 使用翻译后的操作提示
- 添加语言指示器显示当前语言
- 在 `HandleKey()` 中添加 `Ctrl+Space` 处理逻辑(返回 -3)

### 4. 主程序集成 (`cmd/symbol-move/main.go`)

- 导入 i18n 包
- 在 `main()` 中调用 `mgr.LoadConfig()` 加载用户语言偏好
- 在 `waitForSelection()` 中处理语言切换信号 (result == -3)
- 语言切换后自动重新渲染界面

### 5. 测试

创建了完整的单元测试 (`pkg/i18n/i18n_test.go`):
- `TestLanguageToggle` - 测试语言切换功能
- `TestTranslations` - 测试中英文翻译
- `TestEffectNameAndDescription` - 测试特效元数据翻译和回退机制

所有测试通过 ✅

### 6. 文档更新

- 更新 `README.md`:
  * 在"交互式特效选择器"部分添加多语言支持说明
  * 在"键盘控制"部分添加 `Ctrl+Space` 快捷键说明
- 创建 `docs/CONFIG.md` 配置文件说明文档

## 技术亮点

1. **线程安全**: 使用 `sync.RWMutex` 确保并发安全
2. **单例模式**: 全局语言管理器使用单例模式
3. **优雅降级**:
   - 翻译键不存在时返回键本身
   - 英文翻译缺失时回退到中文
   - 配置文件读取失败时使用默认值
4. **持久化**: 用户语言偏好自动保存到 `~/.symbolmove/config.json`
5. **向后兼容**: 扩展 Metadata 结构不影响现有代码

## 用户体验

- ✅ 按 `Ctrl+Space` 即时切换语言
- ✅ 所有界面文本(标题、提示、特效名称、描述)都支持中英文
- ✅ 语言偏好在程序重启后保持
- ✅ 界面显示当前语言指示器
- ✅ 无需手动编辑配置文件

## 文件清单

### 新增文件
- `pkg/i18n/language.go` - 语言管理器
- `pkg/i18n/translations.go` - 翻译映射
- `pkg/i18n/config.go` - 配置管理
- `pkg/i18n/i18n_test.go` - 单元测试
- `docs/CONFIG.md` - 配置文件说明

### 修改文件
- `pkg/effects/effect.go` - 扩展 Metadata 结构
- `pkg/ui/selector/selector.go` - 选择器国际化
- `cmd/symbol-move/main.go` - 主程序集成
- `README.md` - 添加语言切换说明
- 23 个特效的 `effect.go` 文件 - 添加英文元数据

## 编译和运行

```bash
# 编译
go build -o symbol-move.exe ./cmd/symbol-move

# 运行测试
go test ./pkg/i18n -v

# 运行程序
./symbol-move.exe
```

## 已知限制

- 仅支持中英文两种语言
- 语言切换仅在选择器界面有效,特效运行时不支持切换
- 配置文件路径固定为 `~/.symbolmove/config.json`

## 未来扩展建议

1. 添加更多语言支持(日语、韩语等)
2. 支持外部翻译文件(YAML/JSON)
3. 在特效运行时也支持语言切换
4. 添加更多配置选项(主题、字体等)
