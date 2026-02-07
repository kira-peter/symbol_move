package i18n

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config 配置结构
type Config struct {
	Language string `json:"language"`
	Version  string `json:"version"`
}

// getConfigPath 获取配置文件路径
func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".symbolmove")
	configFile := filepath.Join(configDir, "config.json")

	return configFile, nil
}

// LoadConfig 加载配置文件
func (m *Manager) LoadConfig() error {
	configFile, err := getConfigPath()
	if err != nil {
		// 无法获取配置路径，使用默认值
		return nil
	}

	// 读取配置文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		// 配置文件不存在或无法读取，使用默认值
		return nil
	}

	// 解析 JSON
	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		// JSON 格式错误，忽略并使用默认值
		return nil
	}

	// 设置语言
	lang := Language(config.Language)
	if lang == LanguageChinese || lang == LanguageEnglish {
		m.SetLanguage(lang)
	}

	return nil
}

// SaveConfig 保存配置文件
func (m *Manager) SaveConfig() error {
	configFile, err := getConfigPath()
	if err != nil {
		return err
	}

	// 确保配置目录存在
	configDir := filepath.Dir(configFile)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// 构建配置对象
	config := Config{
		Language: string(m.GetCurrent()),
		Version:  "1.0",
	}

	// 序列化为 JSON
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	// 写入文件
	return os.WriteFile(configFile, data, 0644)
}

// ToggleAndSave 切换语言并保存配置
func (m *Manager) ToggleAndSave() (Language, error) {
	lang := m.Toggle()
	err := m.SaveConfig()
	return lang, err
}
