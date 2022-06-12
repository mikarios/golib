package contexts_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/mikarios/golib/contexts"
	"github.com/mikarios/golib/logger"
)

func FuzzCopy(f *testing.F) {
	testKeys := []any{"Key1", "2", string(logger.Settings.TransactionKey)}
	testValues := []any{"1", "Value2", "123"}

	for i := range testKeys {
		f.Add(testKeys[i], testValues[i])
	}

	f.Fuzz(func(t *testing.T, key string, value string) {
		originalCtx := context.WithValue(context.Background(), key, value) // nolint:staticcheck // needed for fuzzing

		copyCtx := contexts.Copy(originalCtx, key)
		if copyCtx.Value(key) != value {
			t.Errorf("invalid value for key %s. Expected: %v, got: %v", key, value, copyCtx.Value(key))
			t.Fail()
		}
	})
}

// nolint:containedctx // this is not important
func TestCopy(t *testing.T) {
	t.Parallel()

	type args struct {
		source context.Context
		keys   []any
	}

	tests := []struct {
		name string
		args args
		want context.Context
	}{
		{
			name: "test copying transaction key",
			args: args{
				source: context.WithValue(context.Background(), logger.Settings.TransactionKey, "tx123"),
				keys:   nil,
			},
			want: context.WithValue(context.Background(), logger.Settings.TransactionKey, "tx123"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := contexts.Copy(tt.args.source, tt.args.keys...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
