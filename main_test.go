package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestBase2ToBase10(t *testing.T) {
	tests := []struct {
		name    string
		num     string
		want    float64
		wantErr error
	}{
		{
			name:    "simple binary number",
			num:     "101101",
			want:    45,
			wantErr: nil,
		},
		{
			name:    "binary number with fraction",
			num:     "101101.101",
			want:    45.625,
			wantErr: nil,
		},
		{
			name:    "binary number with leading zeros",
			num:     "000101101.101",
			want:    45.625,
			wantErr: nil,
		},
		{
			name:    "binary number with trailing zeros",
			num:     "101101.101000",
			want:    45.625,
			wantErr: nil,
		},
		{
			name:    "binary number with all zeros",
			num:     "00000000.00000",
			want:    0,
			wantErr: nil,
		},
		{
			name:    "binary number with all ones",
			num:     "11111111.11111",
			want:    255.96875,
			wantErr: nil,
		},
		{
			name:    "Invalid binary number",
			num:     "101101.1012",
			want:    0,
			wantErr: fmt.Errorf("invalid binary number"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := base2ToBase10(tt.num)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("base2ToBase10(%q) error = %v", tt.num, err)
				return
			}
			if math.Abs(got-tt.want) > 0.0001 {
				t.Errorf("base2ToBase10(%q) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
