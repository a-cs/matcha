package utils

import (
	"backend/internal/defines"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
	"unicode"
)

func ServiceResponse(responseWriter http.ResponseWriter, resp interface{}, status int) {
	responseWriter.WriteHeader(status)
	if resp != nil {
		json.NewEncoder(responseWriter).Encode(resp)
	}
}

func IsSQLInjection(inputs ...string) bool {
	for _, input := range inputs {
		for _, pattern := range defines.CommonSqlPatterns {
			matched, _ := regexp.MatchString(pattern, input)
			if matched {
				return true
			}
		}

		for _, keyword := range defines.SqlKeywords {
			keywordPattern := `\b` + regexp.QuoteMeta(keyword) + `\b`
			matched, _ := regexp.MatchString(keywordPattern, strings.ToLower(input))
			if matched {
				return true
			}
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(defines.EmailRegex)
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return false
	}

	passwordLower := strings.ToLower(password)
	for _, commonWord := range defines.CommonEnglishWords {
		commonWordLower := strings.ToLower(commonWord)
		if passwordLower == commonWordLower {
			return false
		}
		if strings.HasPrefix(passwordLower, commonWordLower) {
			suffix := passwordLower[len(commonWordLower):]
			if len(suffix) > 0 && isOnlyNumbersOrSymbols(suffix) {
				return false
			}
		}
		if strings.HasSuffix(passwordLower, commonWordLower) {
			prefix := passwordLower[:len(passwordLower)-len(commonWordLower)]
			if len(prefix) > 0 && isOnlyNumbersOrSymbols(prefix) {
				return false
			}
		}
	}

	return true
}

func isOnlyNumbersOrSymbols(s string) bool {
	for _, char := range s {
		if !unicode.IsNumber(char) && !unicode.IsPunct(char) && !unicode.IsSymbol(char) {
			return false
		}
	}
	return true
}
