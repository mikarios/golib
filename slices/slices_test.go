package slices

import (
	"reflect"
	"testing"
)

func TestStringContains(t *testing.T) {
	t.Parallel()

	type strArgs struct {
		lst  []string
		elem string
	}

	tests := []struct {
		name string
		args strArgs
		want bool
	}{
		{
			name: "string should contain",
			args: strArgs{
				lst:  []string{"a", "b"},
				elem: "a",
			},
			want: true,
		},
		{
			name: "string should not contain",
			args: strArgs{
				lst:  []string{"a", "b"},
				elem: "c",
			},
			want: false,
		},
		{
			name: "string empty list",
			args: strArgs{
				lst:  nil,
				elem: "c",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Contains(tt.args.lst, tt.args.elem); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntContains(t *testing.T) {
	t.Parallel()

	type args struct {
		lst  []int
		elem int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "int should contain",
			args: args{
				lst:  []int{1, 2},
				elem: 1,
			},
			want: true,
		},
		{
			name: "int should not contain",
			args: args{
				lst:  []int{1, 2},
				elem: 0,
			},
			want: false,
		},
		{
			name: "int empty list",
			args: args{
				lst:  nil,
				elem: 0,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Contains(tt.args.lst, tt.args.elem); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtractStringSlice(t *testing.T) {
	t.Parallel()

	type args struct {
		a []string
		b []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "strings sub 1",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a"},
			},
			want: []string{"b"},
		},
		{
			name: "strings sub 2",
			args: args{
				a: []string{"a", "b"},
				b: []string{"c"},
			},
			want: []string{"a", "b"},
		},
		{
			name: "strings sub 3",
			args: args{
				a: []string{"a", "b"},
				b: nil,
			},
			want: []string{"a", "b"},
		},
		{
			name: "strings sub 4",
			args: args{
				a: nil,
				b: []string{"a"},
			},
			want: []string{},
		},
		{
			name: "strings sub 5",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "b"},
			},
			want: []string{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SubtractSlice(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubtractSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtractIntSlice(t *testing.T) {
	t.Parallel()

	type args struct {
		a []int
		b []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "ints sub 1",
			args: args{
				a: []int{1, 2},
				b: []int{1},
			},
			want: []int{2},
		},
		{
			name: "ints sub 2",
			args: args{
				a: []int{1, 2},
				b: []int{3},
			},
			want: []int{1, 2},
		},
		{
			name: "ints sub 3",
			args: args{
				a: []int{1, 2},
				b: nil,
			},
			want: []int{1, 2},
		},
		{
			name: "ints sub 4",
			args: args{
				a: nil,
				b: []int{1},
			},
			want: []int{},
		},
		{
			name: "ints sub 5",
			args: args{
				a: []int{1, 2},
				b: []int{1, 2},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := SubtractSlice(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubtractSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
