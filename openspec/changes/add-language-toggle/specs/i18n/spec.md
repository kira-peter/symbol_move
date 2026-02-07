## ADDED Requirements

### Requirement: 语言类型定义
系统 SHALL 支持中文(zh)和英文(en)两种语言类型。

#### Scenario: 获取支持的语言列表
- **WHEN** 请求系统支持的语言列表
- **THEN** 返回包含 "zh" 和 "en" 的语言列表

### Requirement: 语言管理器
系统 SHALL 提供一个全局语言管理器,用于管理当前语言状态。

#### Scenario: 获取当前语言
- **WHEN** 查询当前语言
- **THEN** 返回当前设置的语言(默认为中文 "zh")

#### Scenario: 切换语言
- **WHEN** 调用切换语言方法
- **THEN** 当前语言在中英文之间切换
- **AND** 返回切换后的语言

#### Scenario: 设置指定语言
- **WHEN** 设置语言为 "en"
- **THEN** 当前语言变为英文
- **AND** 返回设置成功状态

### Requirement: 翻译文本获取
系统 SHALL 根据当前语言和翻译键返回对应的翻译文本。

#### Scenario: 获取中文翻译
- **WHEN** 当前语言为中文,请求翻译键 "title"
- **THEN** 返回对应的中文文本 "符动世界(SymbolMove)"

#### Scenario: 获取英文翻译
- **WHEN** 当前语言为英文,请求翻译键 "title"
- **THEN** 返回对应的英文文本 "SymbolMove"

#### Scenario: 翻译键不存在
- **WHEN** 请求不存在的翻译键
- **THEN** 返回翻译键本身作为默认值

### Requirement: 配置持久化
系统 SHALL 将用户的语言偏好保存到配置文件中。

#### Scenario: 保存语言配置
- **WHEN** 用户切换语言后
- **THEN** 系统将当前语言写入 `~/.symbolmove/config.json`
- **AND** 配置文件包含 `{"language": "en", "version": "1.0"}` 格式的数据

#### Scenario: 加载语言配置
- **WHEN** 程序启动时
- **THEN** 系统从配置文件读取语言偏好
- **AND** 设置为用户上次选择的语言

#### Scenario: 配置文件不存在
- **WHEN** 配置文件不存在或无法读取
- **THEN** 系统使用默认语言(中文)
- **AND** 不报错,正常启动

#### Scenario: 配置文件格式错误
- **WHEN** 配置文件格式错误或损坏
- **THEN** 系统忽略错误,使用默认语言
- **AND** 继续正常运行

### Requirement: 翻译内容覆盖
系统 SHALL 提供以下界面元素的翻译:

- 主标题 (title)
- 副标题 (subtitle)
- 描述标签 (desc_label)
- 操作提示 (hints)
- 退出提示 (quit_hint)
- 语言指示器文本 (language_indicator)

#### Scenario: 获取主标题翻译
- **WHEN** 当前语言为英文,请求主标题翻译
- **THEN** 返回 "SymbolMove"

#### Scenario: 获取操作提示翻译
- **WHEN** 当前语言为英文,请求操作提示
- **THEN** 返回包含快捷键说明的英文文本
- **AND** 文本包括方向键选择、回车确认、快捷键等说明

### Requirement: 优雅降级
当翻译缺失时,系统 SHALL 提供优雅的降级策略。

#### Scenario: 特效名称缺少英文翻译
- **WHEN** 当前语言为英文,但特效未提供英文名称
- **THEN** 显示中文名称作为后备
- **AND** 不影响其他功能正常运行
