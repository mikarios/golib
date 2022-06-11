package handler_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/mikarios/golib/handler"
)

func TestGetQueryParamInt(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "2"}},
			want:    2,
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "2"}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", 0)
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamInt() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamInt() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamSliceInt(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "1,2"}},
			want:    []int{1, 2},
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "2"}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, ",", []int{})
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamInt() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQueryParamInt() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamSliceInt32(t *testing.T) {
	t.Parallel()

	type args struct {
		params    map[string]string
		key       string
		separator string
	}

	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "1,2"}, separator: ","},
			want:    []int32{1, 2},
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "2"}, separator: ","},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}, separator: ","},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, tt.args.separator, []int32{})
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamInt() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQueryParamInt() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamUInt64(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "2"}},
			want:    2,
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "2"}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", uint64(0))
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamInt() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamInt() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamString(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			want:    "param",
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrParamNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", "")
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamString() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamString() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamSliceString(t *testing.T) {
	t.Parallel()

	type args struct {
		params    map[string]string
		key       string
		separator string
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "param1|param2"}, separator: "|"},
			want:    []string{"param1", "param2"},
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrParamNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, tt.args.separator, []string{})
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamString() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getQueryParamString() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamFloat(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "3.14159"}},
			want:    3.14159,
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "3.14159"}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", float64(0))
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamFloat() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamFloat() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamUUID(t *testing.T) {
	t.Parallel()

	testUUID := uuid.New()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    uuid.UUID
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": testUUID.String()}},
			want:    testUUID,
			wantErr: nil,
		},
		{
			name:    "Valid key and type (without dashes)",
			args:    args{key: "key", params: map[string]string{"key": strings.ReplaceAll(testUUID.String(), "-", "")}},
			want:    testUUID,
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": testUUID.String()}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", uuid.UUID{})
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamUUID() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamUUID() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}

func TestGetQueryParamBool(t *testing.T) {
	t.Parallel()

	type args struct {
		params map[string]string
		key    string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "1"}},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "true"}},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "t"}},
			want:    true,
			wantErr: nil,
		},
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "false"}},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "f"}},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Valid key and type",
			args:    args{key: "key", params: map[string]string{"key": "0"}},
			want:    false,
			wantErr: nil,
		},
		{
			name:    "Invalid key",
			args:    args{key: "key-invalid", params: map[string]string{"key": "1"}},
			wantErr: handler.ErrParamNotFound,
		},
		{
			name:    "Invalid param type (string)",
			args:    args{key: "key", params: map[string]string{"key": "param"}},
			wantErr: handler.ErrConversion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := handler.GetRequestParam(tt.args.params, tt.args.key, "", false)
			if err != nil {
				if (tt.wantErr == nil) || (tt.wantErr != nil && !errors.Is(err, tt.wantErr)) {
					t.Errorf("getQueryParamUUID() = got error = %v, want %v", err, tt.wantErr)
				}

				return
			}

			if got != tt.want {
				t.Errorf("getQueryParamUUID() = got %v, want %v", got, tt.want)
				return
			}
		})
	}
}
