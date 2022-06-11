package stringtools_test

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/mikarios/golib/stringtools"
)

func TestSplitStringByLimit(t *testing.T) {
	t.Parallel()

	type args struct {
		value string
		limit int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "split with remainder",
			args: args{
				value: "asdfagdsfg",
				limit: 3,
			},
			want: []string{"asd", "fag", "dsf", "g"},
		},
		{
			name: "split exactly",
			args: args{
				value: "asdfagdsfg",
				limit: 2,
			},
			want: []string{"as", "df", "ag", "ds", "fg"},
		},
		{
			name: "limit = 0",
			args: args{
				value: "asdfagdsfg",
				limit: 0,
			},
			want: []string{"asdfagdsfg"},
		},
		{
			name: "limit > len(value)",
			args: args{
				value: "asdfagdsfg",
				limit: 100,
			},
			want: []string{"asdfagdsfg"},
		},
		{
			name: "limit > len(value)",
			args: args{
				value: "asdfagdsfg",
				limit: 100,
			},
			want: []string{"asdfagdsfg"},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := stringtools.SplitStringByLimit(tt.args.value, tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitStringByLimit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzSplitStringByLimit(f *testing.F) {
	testStrings := []string{"Hello, world", " ", "!12345"}
	testLimits := []int{1, 2, 3}

	for i := range testStrings {
		f.Add(testStrings[i], testLimits[i])
	}

	f.Fuzz(func(t *testing.T, orig string, lim int) {
		if !utf8.ValidString(orig) {
			return
		}
		split := stringtools.SplitStringByLimit(orig, lim)
		for _, part := range split {
			// Can't use len(part) to take into consideration non utf characters. This way we count how many runes are in each
			// part and not how many bytes it is.
			partLen := 0
			for range part {
				partLen++
			}

			if lim <= 0 {
				if !reflect.DeepEqual(split, []string{orig}) {
					t.Errorf(
						`Since limit is <=0 expected whole string in first element of array. Limit: "%d", orig: "%s", result: "%v"`,
						lim,
						orig,
						split,
					)
					t.Fail()
				}

				lim = partLen
			}

			if partLen > lim {
				t.Errorf(
					`Invalid split. Got more characters (%d) than expected. Orig: "%s", limit: "%d", result: "%v"`,
					len(part),
					orig,
					lim,
					strings.Join(split, ","),
				)
				t.Fail()
			}
		}
	})
}
