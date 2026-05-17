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
		{
			name: "сумма < 100",
			args: args{
				weightedSum: 1,
			},
			want: 1,
		},
		{
			name: "сумма = 100",
			args: args{
				weightedSum: 100,
			},
			want: 0,
		},
		{
			name: "сумма > 100, но остаток < 100",
			args: args{
				weightedSum: 101,
			},
			want: 0,
		},
		{
			name: "сумма > 100, остаток == 100",
			args: args{
				weightedSum: 201,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateControlNumber(tt.args.weightedSum); got != tt.want {
				t.Errorf("calculateControlNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseControlNumber(t *testing.T) {
	type args struct {
		controlPart string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "валидное число",
			args: args{
				controlPart: "95",
			},
			want: 95,
		},
		{
			name: "число с незначащими пробелами",
			args: args{
				controlPart: "001",
			},
			want: 1,
		},
		{
			name: "отрицательное число",
			args: args{
				controlPart: "-101",
			},
			want: -101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseControlNumber(tt.args.controlPart); got != tt.want {
				t.Errorf("parseControlNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
