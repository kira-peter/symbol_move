## MODIFIED Requirements

### Requirement: 特效元数据结构
特效元数据 SHALL 支持多语言名称和描述。

#### Scenario: 元数据包含英文字段
- **WHEN** 定义特效元数据
- **THEN** 元数据结构包含 `NameEN` 字段用于英文名称
- **AND** 元数据结构包含 `DescriptionEN` 字段用于英文描述
- **AND** 保留原有的 `Name` 和 `Description` 字段用于中文

#### Scenario: 创建包含双语的特效元数据
- **WHEN** 创建特效实例并返回元数据
- **THEN** 元数据同时包含中文和英文的名称
- **AND** 元数据同时包含中文和英文的描述
- **AND** 中英文内容准确对应,语义一致

### Requirement: 特效元数据向后兼容
特效元数据的扩展 SHALL 保持向后兼容。

#### Scenario: 未提供英文字段
- **WHEN** 特效元数据未填写 `NameEN` 或 `DescriptionEN`
- **THEN** 这些字段为空字符串
- **AND** 不影响特效的正常注册和运行
- **AND** 界面可以降级使用中文显示

#### Scenario: 只提供中文元数据的旧特效
- **WHEN** 旧的特效只提供了 `Name` 和 `Description`
- **THEN** 特效仍然可以正常注册
- **AND** 特效在列表中正常显示(使用中文)
- **AND** 不产生运行时错误

## ADDED Requirements

### Requirement: 特效元数据最佳实践
所有新增特效 SHALL 提供完整的双语元数据。

#### Scenario: 新特效提供双语元数据
- **WHEN** 创建新的特效
- **THEN** 元数据包含有意义的英文名称
- **AND** 元数据包含准确的英文描述
- **AND** 英文描述简洁明了,准确传达特效功能

### Requirement: 元数据翻译质量
英文元数据 SHALL 准确反映特效的功能和特点。

#### Scenario: 特效名称翻译准确性
- **WHEN** 特效名称为 "矩阵字符雨"
- **THEN** 英文名称为 "Matrix Rain" 或类似的准确翻译
- **AND** 名称简洁,符合英文命名习惯

#### Scenario: 特效描述翻译准确性
- **WHEN** 特效描述为 "经典黑客帝国风格字符雨效果"
- **THEN** 英文描述准确传达相同含义
- **AND** 描述语法正确,专业易懂
