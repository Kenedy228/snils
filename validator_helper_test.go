package snils

import (
	"strings"
	"testing"
)

func Test_validateLength(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "корректная длина",
			args: args{
				snils: strings.Repeat("a", SNILSLength),
			},
			wantErr: false,
		},
		{
			name: "длина меньшше",
			args: args{
				snils: strings.Repeat("a", SNILSLength-1),
			},
			wantErr: true,
		},
		{
			name: "длина больше",
			args: args{
				snils: strings.Repeat("a", SNILSLength+1),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateLength(tt.args.snils); (err != nil) != tt.wantErr {
				t.Errorf("validateLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateDigitsOnly(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "все цифры",
			args: args{
				snils: strings.Repeat("1", 10),
			},
			wantErr: false,
		},
		{
			name: "с буквами",
			args: args{
				snils: "123aa",
			},
			wantErr: true,
		},
		{
			name: "с пробелами",
			args: args{
				snils: "123 321",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDigitsOnly(tt.args.snils); (err != nil) != tt.wantErr {
				t.Errorf("validateDigitsOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateForbiddenSNILS(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "запрещенный снилс",
			args: args{
				snils: ForbiddenSNILS,
			},
			wantErr: true,
		},
		{
			name: "незапрещенный снилс",
			args: args{
				snils: strings.Repeat("1", 11),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateForbiddenSNILS(tt.args.snils); (err != nil) != tt.wantErr {
				t.Errorf("validateForbiddenSNILS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_shouldValidateChecksum(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "минимум",
			args: args{
				snils: "001001998",
			},
			want: false,
		},
		{
			name: "меньше минимума",
			args: args{
				snils: "001001997",
			},
			want: false,
		},
		{
			name: "больше минимума",
			args: args{
				snils: "001001999",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shouldValidateChecksum(tt.args.snils); got != tt.want {
				t.Errorf("shouldValidateChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
