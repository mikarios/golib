package dateformattranslator_test

import (
	"testing"
	"time"

	"github.com/mikarios/golib/dates/dateformattranslator"
)

// nolint: funlen // well... I wanted to cover all cases
func TestConvertDateFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		format string
		want   string
	}{
		{
			format: "M",
			want:   "1",
		},
		{
			format: "MM",
			want:   "01",
		},
		{
			format: "MMM",
			want:   "Jan",
		},
		{
			format: "MMMM",
			want:   "January",
		},
		{
			format: "dddd",
			want:   "Monday",
		},
		{
			format: "ddd",
			want:   "Mon",
		},
		{
			format: "dd",
			want:   "Mon",
		},
		{
			format: "DD",
			want:   "02",
		},
		{
			format: "D",
			want:   "2",
		},
		{
			format: "YYYY",
			want:   "2006",
		},
		{
			format: "YY",
			want:   "06",
		},
		{
			format: "yyyy",
			want:   "2006",
		},
		{
			format: "yy",
			want:   "06",
		},
		{
			format: "A",
			want:   "PM",
		},
		{
			format: "a",
			want:   "pm",
		},
		{
			format: "H",
			want:   "15",
		},
		{
			format: "HH",
			want:   "15",
		},
		{
			format: "h",
			want:   "3",
		},
		{
			format: "hh",
			want:   "03",
		},
		{
			format: "mm",
			want:   "04",
		},
		{
			format: "m",
			want:   "4",
		},
		{
			format: "s",
			want:   "5",
		},
		{
			format: "ss",
			want:   "05",
		},
		{
			format: "S",
			want:   "0",
		},
		{
			format: "SSSSSS",
			want:   "000000",
		},
		{
			format: "z",
			want:   "MST",
		},
		{
			format: "zz",
			want:   "MST",
		},
		{
			format: "Z",
			want:   "-07:00",
		},
		{
			format: "ZZ",
			want:   "-0700",
		},
		{
			format: "ddd MMM D HH:mm:ss YYYY",
			want:   "Mon Jan 2 15:04:05 2006",
		},
		{
			format: "ddd MMM D HH:mm:ss YYYY",
			want:   "Mon Jan 2 15:04:05 2006",
		},
		{
			format: "dd MMM D HH:mm:ss z YYYY",
			want:   "Mon Jan 2 15:04:05 MST 2006",
		},
		{
			format: "dd MMM D HH:mm:ss zz YYYY",
			want:   "Mon Jan 2 15:04:05 MST 2006",
		},
		{
			format: "dd MMM DD HH:mm:ss ZZ YYYY",
			want:   time.RubyDate,
		},
		{
			format: "dd MMM DD HH:mm:ss ZZ YYYY",
			want:   time.RubyDate,
		},
		{
			format: "DD MMM YY HH:mm z",
			want:   time.RFC822,
		},
		{
			format: "DD MMM YY HH:mm ZZ",
			want:   time.RFC822Z,
		},
		{
			format: "dddd, DD-MMM-YY HH:mm:ss z",
			want:   time.RFC850,
		},
		{
			format: "ddd, DD MMM YYYY HH:mm:ss z",
			want:   time.RFC1123,
		},
		{
			format: "ddd, DD MMM YYYY HH:mm:ss ZZ",
			want:   time.RFC1123Z,
		},
		{
			format: "h:mmA",
			want:   time.Kitchen,
		},
		{
			format: "MMM D HH:mm:ss",
			want:   "Jan 2 15:04:05",
		},
		{
			format: "MMM D HH:mm:ss.SSS",
			want:   "Jan 2 15:04:05.000",
		},
		{
			format: "MMM D HH:mm:ss.SSSSSS",
			want:   "Jan 2 15:04:05.000000",
		},
		{
			format: "MMM D HH:mm:ss.SSSSSSSSS",
			want:   "Jan 2 15:04:05.000000000",
		},
		{
			format: "D/M/YYYY",
			want:   "2/1/2006",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.format, func(t *testing.T) {
			t.Parallel()

			if got := dateformattranslator.Translate(tt.format); got != tt.want {
				t.Errorf("ConvertDateFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
