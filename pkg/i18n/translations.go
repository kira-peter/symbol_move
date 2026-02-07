package i18n

// 翻译键常量
const (
	KeyTitle            = "title"
	KeySubtitle         = "subtitle"
	KeyDescLabel        = "desc_label"
	KeyHints            = "hints"
	KeyLanguageIndicator = "lang_indicator"
)

// uiTexts 界面文本翻译映射
var uiTexts = map[Language]map[string]string{
	LanguageChinese: {
		KeyTitle:             "符动世界(SymbolMove)",
		KeySubtitle:          "字符符号在动，创造世界",
		KeyDescLabel:         "描述:",
		KeyHints:             "↑↓←→:选择 | Enter:确认 | 1-9/0:快捷键 | Ctrl+Space:切换语言 | q/Ctrl+C:退出",
		KeyLanguageIndicator: "中文",
	},
	LanguageEnglish: {
		KeyTitle:             "SymbolMove",
		KeySubtitle:          "Characters in Motion, Creating Worlds",
		KeyDescLabel:         "Description:",
		KeyHints:             "↑↓←→:Select | Enter:Confirm | 1-9/0:Shortcut | Ctrl+Space:Switch Lang | q/Ctrl+C:Quit",
		KeyLanguageIndicator: "English",
	},
}

// T 获取翻译文本（便捷函数）
func T(key string) string {
	return GetManager().T(key)
}

// T 获取指定键的翻译文本
func (m *Manager) T(key string) string {
	m.mu.RLock()
	lang := m.current
	m.mu.RUnlock()

	// 获取当前语言的翻译映射
	translations, ok := uiTexts[lang]
	if !ok {
		// 语言不存在，使用中文作为后备
		translations = uiTexts[LanguageChinese]
	}

	// 获取翻译文本
	text, ok := translations[key]
	if !ok {
		// 翻译键不存在，返回键本身
		return key
	}

	return text
}

// GetEffectName 获取特效名称（根据当前语言）
func (m *Manager) GetEffectName(nameCN, nameEN string) string {
	if m.IsEnglish() && nameEN != "" {
		return nameEN
	}
	return nameCN
}

// GetEffectDescription 获取特效描述（根据当前语言）
func (m *Manager) GetEffectDescription(descCN, descEN string) string {
	if m.IsEnglish() && descEN != "" {
		return descEN
	}
	return descCN
}
