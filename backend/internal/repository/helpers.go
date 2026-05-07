package repository

import (
	"strconv"
	"strings"
	"time"
)

// parseFloat извлекает число из строки типа "35.000 р."
func parseFloat(s string) float64 {
	s = strings.ReplaceAll(s, " р.", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return val
}

// formatMoney форматирует число в строку типа "35.000 р."
func formatMoney(val float64) string {
	if val == 0 {
		return "0 р."
	}
	intPart := int(val)
	s := strconv.Itoa(intPart)
	if len(s) > 3 {
		s = s[:len(s)-3] + "." + s[len(s)-3:]
	}
	return s + " р."
}

// daysUntil возвращает количество дней до указанной даты
func daysUntil(dateStr string) int {
	if dateStr == "" {
		return 0
	}

	var t time.Time
	var err error

	if strings.Contains(dateStr, "-") {
		t, err = time.Parse("2006-01-02", dateStr)
	} else {
		t, err = time.Parse("02.01.2006", dateStr)
	}

	if err != nil {
		return 0
	}

	days := int(time.Until(t).Hours() / 24)
	return days
}
