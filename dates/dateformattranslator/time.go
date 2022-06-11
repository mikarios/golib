package dateformattranslator

import (
	"strings"
)

// Translate gets a format as is defined in momentjs and returns the equivalent in go.
func Translate(format string) string {
	replacer := strings.NewReplacer(
		"MMMM", "January",
		"MMM", "Jan",
		"MM", "01",
		"M", "1",

		"dddd", "Monday",
		"ddd", "Mon",
		"dd", "Mon",
		"DD", "02",
		"D", "2",

		"YYYY", "2006",
		"yyyy", "2006",
		"YY", "06",
		"yy", "06",

		"A", "PM",
		"a", "pm",

		"HH", "15",
		"H", "15",
		"hh", "03",
		"h", "3",

		"mm", "04",
		"m", "4",

		"ss", "05",
		"s", "5",

		"S", "0",

		"zz", "MST",
		"z", "MST",

		"ZZ", "-0700",
		"Z", "-07:00",
	)

	return replacer.Replace(format)
}
