package snils

import "testing"

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
			name: "ответ 95",
			args: args{
				numberPart: "112233445",
			},
			want: 95,
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
