package languagehelper

import (
	"strings"
)

// Retrieve is the default implementation
func Retrieve(acceptedLangagesHeader string) string {
	const en = "en"
	const de = "de"
	deIndex := strings.Index(acceptedLangagesHeader, de)
	enIndex := strings.Index(acceptedLangagesHeader, en)
	if deIndex == -1 {
		return en
	}
	if enIndex != -1 {
		if deIndex < enIndex {
			return de
		}
		return en
	}
	return de
}
