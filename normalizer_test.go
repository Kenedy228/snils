package snils

import (
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "только цифры",
			args: args{
				snils: strings.Repeat("1", 30),
			},
			want: strings.Repeat("1", 30),
		},
		{
			name: "цифры и буквы",
			args: args{
				snils: "123abcd",
			},
			want: "123abcd",
		},
		{
			name: "цифры и пробелы",
			args: args{
				snils: "123 321   4555",
			},
			want: "1233214555",
		},
		{
			name: "цифры и дефисы",
			args: args{
				snils: "123-321-455",
			},
			want: "123321455",
		},
		{
			name: "цифры, дефисы и пробелы",
			args: args{
				snils: " 123-321 455--",
			},
			want: "123321455",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.args.snils); got != tt.want {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
