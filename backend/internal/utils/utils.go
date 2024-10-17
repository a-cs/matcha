package utils

import (
	"backend/internal/defines"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

func ServiceResponse(responseWriter http.ResponseWriter, resp interface{}, status int) {
	responseWriter.WriteHeader(status)
	if resp != nil {
		json.NewEncoder(responseWriter).Encode(resp)
	}
}

func IsSQLInjection(inputs ...string) bool {
	for _, input := range inputs {
		// Check for common SQL injection patterns
		for _, pattern := range defines.CommonSqlPatterns {
			matched, _ := regexp.MatchString(pattern, input)
			if matched {
				return true
			}
		}

		// Check for common SQL keywords
		for _, keyword := range defines.SqlKeywords {
			if strings.Contains(strings.ToLower(input), keyword) {
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
