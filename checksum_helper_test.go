package snils

import (
	"strings"
	"testing"
)

func Test_calculateWeightedSum(t *testing.T) {
	type args struct {
		numberPart string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "случайный набор",
			args: args{
				numberPart: "112233445",
			},
			want: 95,
		},
		{
			name: "все 0",
			args: args{
				numberPart: strings.Repeat("0", 9),
			},
			want: 0,
		},
		{
			name: "все 1",
			args: args{
				numberPart: strings.Repeat("1", 9),
			},
			want: 45,
		},
		{
			name: "арифметическая прогрессия",
			args: args{
				numberPart: "123456789",
			},
			want: 165,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateWeightedSum(tt.args.numberPart); got != tt.want {
				t.Errorf("calculateWeightedSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateControlNumber(t *testing.T) {
	type args struct {
		weightedSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateControlNumber(tt.args.weightedSum); got != tt.want {
				t.Errorf("calculateControlNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
