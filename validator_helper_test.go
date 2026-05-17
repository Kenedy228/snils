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
				snils: "00100199899",
			},
			want: false,
		},
		{
			name: "меньше минимума",
			args: args{
				snils: "00100199799",
			},
			want: false,
		},
		{
			name: "больше минимума",
			args: args{
				snils: "00100199999",
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

func Test_validateChecksum(t *testing.T) {
	type args struct {
		snils string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "СНИЛС из примера ПФ - верный",
			args: args{
				snils: "11223344595",
			},
			wantErr: false,
		},
		{
			name: "СНИЛС из примера ПФ - неверный",
			args: args{
				snils: "11223344594",
			},
			wantErr: true,
		},
		{
			name: "контрольное число 99, контрольная сумма 99",
			args: args{
				snils: "00101989999",
			},
			wantErr: false,
		},
		{
			name: "контрольное число 00, контрольная сумма 100",
			args: args{
				snils: "00101998900",
			},
			wantErr: false,
		},
		{
			name: "контрольное число 00, контрольная сумма 101",
			args: args{
				snils: "00101999800",
			},
			wantErr: false,
		},
		{
			name: "контрольное число 01, контрольная сумма 102",
			args: args{
				snils: "00101999901",
			},
			wantErr: false,
		},
		{
			name: "контрольное число 64, контрольная сумма 165",
			args: args{
				snils: "12345678964",
			},
			wantErr: false,
		},
		{
			name: "контрольное число 65, контрольная сумма 165",
			args: args{
				snils: "12345678965",
			},
			wantErr: true,
		},
		{
			name: "контрольное число 44, контрольная сумма 142",
			args: args{
				snils: "90114404442",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateChecksum(tt.args.snils); (err != nil) != tt.wantErr {
				t.Errorf("validateChecksum() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
