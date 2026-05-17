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

func Test_validateContent(t *testing.T) {
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
			if err := validateContent(tt.args.snils); (err != nil) != tt.wantErr {
				t.Errorf("validateContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
