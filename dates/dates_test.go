package dates_test

import (
	"testing"
	"time"

	"github.com/mikarios/golib/dates"
)

func TestMatchesDay(t *testing.T) {
	t.Parallel()

	type args struct {
		now   time.Time
		check time.Time
	}

	now := time.Now()

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "now",
			args: args{
				now:   now,
				check: now,
			},
			want: true,
		},
		{
			name: "now not matches",
			args: args{
				now:   now,
				check: now.Add(24 * time.Hour),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := dates.MatchesDay(tt.args.now, tt.args.check); got != tt.want {
				t.Errorf("MatchesDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
