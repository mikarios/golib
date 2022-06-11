package contexts_test

import (
	"context"
	"testing"

	"github.com/mikarios/golib/contexts"
)

func FuzzCopy(f *testing.F) {
	testKeys := []any{"Key1", "2"}
	testValues := []any{"1", "Value2"}

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
