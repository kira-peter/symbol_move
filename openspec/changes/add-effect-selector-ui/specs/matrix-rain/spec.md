# Matrix Rain Effect Specification Delta

## MODIFIED Requirements

### Requirement: 字符雨基础渲染
系统 SHALL 在终端中渲染连续下落的字符流，创造类似"黑客帝国"的视觉效果，并支持通过标准接口被主程序调用。

#### Scenario: 基础字符雨显示
- **WHEN** 程序启动或通过接口调用
- **THEN** 终端填充多条从上到下移动的字符流
- **AND** 字符流以绿色显示
- **AND** 每条字符流独立移动

#### Scenario: 字符随机性
- **WHEN** 字符流下落时
- **THEN** 每个位置的字符随机从字符集中选择
- **AND** 字符流的长度随机变化
- **AND** 新字符流的出现位置随机分布

#### Scenario: 通过接口调用
- **WHEN** 主程序通过 Effect 接口启动字符雨
- **THEN** 字符雨效果正确初始化并运行
- **AND** 接收传入的 tcell.Screen 实例
- **AND** 支持外部控制生命周期

## ADDED Requirements

### Requirement: 特效接口实现
字符雨特效 SHALL 实现标准的 Effect 接口，以便被主程序管理。

#### Scenario: 实现 Effect 接口
- **WHEN** 字符雨特效被实例化时
- **THEN** 提供 Init(screen) 方法初始化特效
- **AND** 提供 Run() 方法运行特效主循环
- **AND** 提供 Cleanup() 方法清理资源
- **AND** 提供 Metadata() 方法返回特效元数据

#### Scenario: 元数据提供
- **WHEN** 主程序查询特效元数据时
- **THEN** 返回特效名称"矩阵字符雨"
- **AND** 返回特效描述
- **AND** 返回特效版本信息
- **AND** 可选返回作者和标签信息

#### Scenario: 生命周期管理
- **WHEN** 特效从 Init → Run → Cleanup 状态转换时
- **THEN** 每个阶段正确执行
- **AND** 资源按需分配和释放
- **AND** 支持中途中断（接收 quit 信号）

### Requirement: 独立可执行性
字符雨特效 SHALL 既可以作为独立程序运行，也可以作为模块被调用。

#### Scenario: 独立命令行运行
- **WHEN** 通过 `cmd/matrix-rain/main.go` 启动时
- **THEN** 作为独立程序运行
- **AND** 支持所有命令行参数
- **AND** 直接控制终端

#### Scenario: 模块化调用
- **WHEN** 通过主程序接口调用时
- **THEN** 作为模块运行
- **AND** 使用传入的 screen 实例
- **AND** 遵守接口约定
