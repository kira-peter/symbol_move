package i18n

import (
	"testing"
)

func TestLanguageToggle(t *testing.T) {
	mgr := GetManager()

	// 默认应该是中文
	if mgr.GetCurrent() != LanguageChinese {
		t.Errorf("Expected default language to be Chinese, got %s", mgr.GetCurrent())
	}

	// 切换到英文
	newLang := mgr.Toggle()
	if newLang != LanguageEnglish {
		t.Errorf("Expected language to be English after toggle, got %s", newLang)
	}

	// 再次切换回中文
	newLang = mgr.Toggle()
	if newLang != LanguageChinese {
		t.Errorf("Expected language to be Chinese after second toggle, got %s", newLang)
	}
}

func TestTranslations(t *testing.T) {
	mgr := GetManager()

	// 测试中文翻译
	mgr.SetLanguage(LanguageChinese)
	title := mgr.T(KeyTitle)
	if title != "符动世界(SymbolMove)" {
		t.Errorf("Expected Chinese title, got %s", title)
	}

	// 测试英文翻译
	mgr.SetLanguage(LanguageEnglish)
	title = mgr.T(KeyTitle)
	if title != "SymbolMove" {
		t.Errorf("Expected English title, got %s", title)
	}
}

func TestEffectNameAndDescription(t *testing.T) {
	mgr := GetManager()

	nameCN := "矩阵数字雨"
	nameEN := "Matrix Rain"
	descCN := "经典黑客帝国风格的数字雨特效"
	descEN := "Classic Matrix-style digital rain effect"

	// 测试中文
	mgr.SetLanguage(LanguageChinese)
	if name := mgr.GetEffectName(nameCN, nameEN); name != nameCN {
		t.Errorf("Expected Chinese name, got %s", name)
	}
	if desc := mgr.GetEffectDescription(descCN, descEN); desc != descCN {
		t.Errorf("Expected Chinese description, got %s", desc)
	}

	// 测试英文
	mgr.SetLanguage(LanguageEnglish)
	if name := mgr.GetEffectName(nameCN, nameEN); name != nameEN {
		t.Errorf("Expected English name, got %s", name)
	}
	if desc := mgr.GetEffectDescription(descCN, descEN); desc != descEN {
		t.Errorf("Expected English description, got %s", desc)
	}

	// 测试英文缺失时的后备
	mgr.SetLanguage(LanguageEnglish)
	if name := mgr.GetEffectName(nameCN, ""); name != nameCN {
		t.Errorf("Expected fallback to Chinese name when English is empty, got %s", name)
	}
}
