package snils

import "testing"

func Test_isDelimiter(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "разделитель",
			args: args{
				r: delimiters[0],
			},
			want: true,
		},
		{
			name: "не разделитель",
			args: args{
				r: 'a',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDelimiter(tt.args.r); got != tt.want {
				t.Errorf("isDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDigit(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "цифра",
			args: args{
				r: '1',
			},
			want: true,
		},
		{
			name: "граничная цифра 0",
			args: args{
				r: '0',
			},
			want: true,
		},
		{
			name: "граничная цифра 9",
			args: args{
				r: '9',
			},
			want: true,
		},
		{
			name: "не цифра",
			args: args{
				r: 'a',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.r); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSpace(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "пробел",
			args: args{
				r: ' ',
			},
			want: true,
		},
		{
			name: "не пробел",
			args: args{
				r: 'a',
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSpace(tt.args.r); got != tt.want {
				t.Errorf("isSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
