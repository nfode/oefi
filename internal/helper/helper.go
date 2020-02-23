package helper

import (
	"time"
)

func FormatTimeStr(str string) *string {
	result, _ := time.Parse(time.RFC3339, str)
	formatedStr := result.Format(time.Kitchen)
	return &formatedStr
}

