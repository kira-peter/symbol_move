package effects

import (
	"fmt"
	"sort"
	"sync"
)

// EffectFactory 特效工厂函数类型
// 返回一个新的特效实例
type EffectFactory func() Effect

// Registry 特效注册表
type Registry struct {
	mu        sync.RWMutex
	effects   map[string]EffectFactory
	order     []string // 保持注册顺序
}

// NewRegistry 创建新的特效注册表
func NewRegistry() *Registry {
	return &Registry{
		effects: make(map[string]EffectFactory),
		order:   make([]string, 0),
	}
}

// Register 注册一个特效
func (r *Registry) Register(factory EffectFactory) error {
	// 创建临时实例以获取元数据
	effect := factory()
	metadata := effect.Metadata()

	if metadata.ID == "" {
		return fmt.Errorf("特效 ID 不能为空")
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// 检查是否已注册
	if _, exists := r.effects[metadata.ID]; exists {
		return fmt.Errorf("特效 '%s' 已经注册", metadata.ID)
	}

	r.effects[metadata.ID] = factory
	r.order = append(r.order, metadata.ID)

	return nil
}

// Get 根据 ID 获取特效工厂函数
func (r *Registry) Get(id string) (EffectFactory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	factory, exists := r.effects[id]
	if !exists {
		return nil, fmt.Errorf("未找到特效: %s", id)
	}

	return factory, nil
}

// List 返回所有已注册特效的元数据列表（按注册顺序）
func (r *Registry) List() []Metadata {
	r.mu.RLock()
	defer r.mu.RUnlock()

	metadataList := make([]Metadata, 0, len(r.order))

	for _, id := range r.order {
		factory := r.effects[id]
		effect := factory()
		metadataList = append(metadataList, effect.Metadata())
	}

	return metadataList
}

// ListSorted 返回排序后的特效元数据列表
func (r *Registry) ListSorted() []Metadata {
	list := r.List()

	sort.Slice(list, func(i, j int) bool {
		return list[i].Name < list[j].Name
	})

	return list
}

// Count 返回已注册特效的数量
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.effects)
}

// Has 检查特效是否已注册
func (r *Registry) Has(id string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.effects[id]
	return exists
}

// GlobalRegistry 全局特效注册表
var GlobalRegistry = NewRegistry()

// Register 注册特效到全局注册表（便捷函数）
func Register(factory EffectFactory) error {
	return GlobalRegistry.Register(factory)
}

// Get 从全局注册表获取特效（便捷函数）
func Get(id string) (EffectFactory, error) {
	return GlobalRegistry.Get(id)
}

// List 列出全局注册表中的所有特效（便捷函数）
func List() []Metadata {
	return GlobalRegistry.List()
}
