## ADDED Requirements

### Requirement: 语言切换快捷键
选择器界面 SHALL 支持 `Ctrl+Space` 快捷键切换语言。

#### Scenario: 按下 Ctrl+Space 切换到英文
- **WHEN** 当前语言为中文,用户按下 `Ctrl+Space`
- **THEN** 界面语言切换为英文
- **AND** 所有界面文本更新为英文显示
- **AND** 语言偏好保存到配置文件

#### Scenario: 按下 Ctrl+Space 切换到中文
- **WHEN** 当前语言为英文,用户按下 `Ctrl+Space`
- **THEN** 界面语言切换为中文
- **AND** 所有界面文本更新为中文显示
- **AND** 语言偏好保存到配置文件

#### Scenario: 切换后立即重新渲染
- **WHEN** 用户切换语言
- **THEN** 选择器界面立即重新渲染
- **AND** 用户无需任何额外操作即可看到语言变化

### Requirement: 语言指示器显示
选择器界面 SHALL 在标题区域显示当前语言指示器。

#### Scenario: 显示中文语言指示器
- **WHEN** 当前语言为中文
- **THEN** 在标题附近显示 "语言: 中文" 或 "中文/English"
- **AND** 当前语言以高亮或特殊样式显示

#### Scenario: 显示英文语言指示器
- **WHEN** 当前语言为英文
- **THEN** 在标题附近显示 "Language: English" 或 "中文/English"
- **AND** 当前语言以高亮或特殊样式显示

### Requirement: 界面文本国际化
选择器界面的所有文本 SHALL 根据当前语言显示。

#### Scenario: 英文界面显示
- **WHEN** 当前语言为英文
- **THEN** 主标题显示 "SymbolMove"
- **AND** 副标题显示 "Characters in Motion, Creating Worlds"
- **AND** 描述标签显示 "Description:"
- **AND** 操作提示显示英文说明

#### Scenario: 中文界面显示
- **WHEN** 当前语言为中文
- **THEN** 主标题显示 "符动世界(SymbolMove)"
- **AND** 副标题显示 "字符符号在动,创造世界"
- **AND** 描述标签显示 "描述:"
- **AND** 操作提示显示中文说明

### Requirement: 特效名称多语言显示
选择器 SHALL 根据当前语言显示特效名称。

#### Scenario: 显示英文特效名称
- **WHEN** 当前语言为英文,且特效提供了英文名称
- **THEN** 特效列表显示英文名称
- **AND** 特效描述显示英文描述

#### Scenario: 特效缺少英文翻译时降级
- **WHEN** 当前语言为英文,但特效未提供英文名称
- **THEN** 特效列表显示中文名称
- **AND** 特效描述显示中文描述

### Requirement: 操作提示更新
选择器的操作提示 SHALL 包含语言切换快捷键说明。

#### Scenario: 中文操作提示
- **WHEN** 当前语言为中文
- **THEN** 操作提示包含 "Ctrl+Space:切换语言"
- **AND** 操作提示显示完整的中文快捷键说明

#### Scenario: 英文操作提示
- **WHEN** 当前语言为英文
- **THEN** 操作提示包含 "Ctrl+Space:Switch Language"
- **AND** 操作提示显示完整的英文快捷键说明

## MODIFIED Requirements

### Requirement: 选择器键盘事件处理
选择器的键盘事件处理 SHALL 支持 `Ctrl+Space` 快捷键。

#### Scenario: 处理 Ctrl+Space 事件
- **WHEN** 用户在选择器界面按下 `Ctrl+Space`
- **THEN** 系统调用语言切换方法
- **AND** 触发界面重新渲染
- **AND** 不影响其他快捷键功能

#### Scenario: Ctrl+Space 不与其他快捷键冲突
- **WHEN** 用户使用其他快捷键(↑↓←→, Enter, q, 1-9/0 等)
- **THEN** 这些快捷键功能正常工作
- **AND** 不受语言切换功能影响
