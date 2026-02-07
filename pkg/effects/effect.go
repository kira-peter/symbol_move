package effects

import "github.com/gdamore/tcell/v2"

// Effect 定义特效的标准接口
// 所有特效都必须实现此接口以便被主程序管理
type Effect interface {
	// Metadata 返回特效的元数据
	Metadata() Metadata

	// Init 初始化特效
	// screen: 终端屏幕实例
	// 返回错误时表示初始化失败
	Init(screen tcell.Screen) error

	// Run 运行特效主循环
	// quit: 接收退出信号的通道，收到信号时应立即停止
	// 返回错误时表示运行时出错
	Run(quit <-chan struct{}) error

	// Cleanup 清理特效资源
	// 在特效结束时调用，确保资源被正确释放
	Cleanup() error
}

// Metadata 特效元数据
type Metadata struct {
	// ID 特效唯一标识符（kebab-case）
	ID string

	// Name 特效显示名称（中文）
	Name string

	// Description 特效描述（简短说明，中文）
	Description string

	// NameEN 特效显示名称（英文）
	NameEN string

	// DescriptionEN 特效描述（简短说明，英文）
	DescriptionEN string

	// LongDescription 详细描述（可选）
	LongDescription string

	// Author 作者信息（可选）
	Author string

	// Version 版本号（可选）
	Version string

	// Tags 标签列表（可选）
	Tags []string
}
