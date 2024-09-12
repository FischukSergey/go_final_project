package repeatrule

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNextDate(t *testing.T) {
	type args struct {
		now    time.Time
		date   string
		repeat string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "test d 7",
			args: args{
				now:    time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC),
				date:   "20240113",
				repeat: "d 7",
			},
			want:    "20240127",
			wantErr: nil,
		},
		{
			name: "test y",
			args: args{
				now:    time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC),
				date:   "20240229",
				repeat: "y",
			},
			want:    "20250301",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextDate(tt.args.now, tt.args.date, tt.args.repeat)
			
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}
