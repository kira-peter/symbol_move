package i18n

import (
	"sync"
)

// Language 语言类型
type Language string

const (
	// LanguageChinese 中文
	LanguageChinese Language = "zh"
	// LanguageEnglish 英文
	LanguageEnglish Language = "en"
)

// Manager 语言管理器
type Manager struct {
	mu      sync.RWMutex
	current Language
}

// globalManager 全局语言管理器实例
var globalManager *Manager
var once sync.Once

// GetManager 获取全局语言管理器实例
func GetManager() *Manager {
	once.Do(func() {
		globalManager = &Manager{
			current: LanguageChinese, // 默认中文
		}
	})
	return globalManager
}

// GetCurrent 获取当前语言
func (m *Manager) GetCurrent() Language {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.current
}

// SetLanguage 设置语言
func (m *Manager) SetLanguage(lang Language) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.current = lang
}

// Toggle 切换语言（中英文之间）
func (m *Manager) Toggle() Language {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.current == LanguageChinese {
		m.current = LanguageEnglish
	} else {
		m.current = LanguageChinese
	}

	return m.current
}

// IsEnglish 判断当前是否为英文
func (m *Manager) IsEnglish() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.current == LanguageEnglish
}

// IsChinese 判断当前是否为中文
func (m *Manager) IsChinese() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.current == LanguageChinese
}
