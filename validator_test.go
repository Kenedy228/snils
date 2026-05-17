package snils

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		normalizedSnils string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "пример ПФ",
			args: args{
				normalizedSnils: "11223344595",
			},
			wantErr: false,
		},
		{
			name: "контрольная сумма < 100",
			args: args{
				normalizedSnils: "00101989999",
			},
			wantErr: false,
		},
		{
			name: "контрольная сумма = 100",
			args: args{
				normalizedSnils: "00101998900",
			},
			wantErr: false,
		},
		{
			name: "контрольная сумма > 100, остаток 0",
			args: args{
				normalizedSnils: "00101999800",
			},
			wantErr: false,
		},
		{
			name: "контрольная сумма > 100, остаток 1",
			args: args{
				normalizedSnils: "00101999901",
			},
			wantErr: false,
		},
		{
			name: "длина < 11 цифр",
			args: args{
				normalizedSnils: "1122334459",
			},
			wantErr: true,
		},
		{
			name: "длина > 11 цифр",
			args: args{
				normalizedSnils: "112233445950",
			},
			wantErr: true,
		},
		{
			name: "содержит недопустимые символы",
			args: args{
				normalizedSnils: "1122334459a",
			},
			wantErr: true,
		},
		{
			name: "запрещенный СНИЛС",
			args: args{
				normalizedSnils: "00000000000",
			},
			wantErr: true,
		},
		{
			name: "контрольное число != контрольная сумма",
			args: args{
				normalizedSnils: "11223344594",
			},
			wantErr: true,
		},
		{
			name: "валидация контрольной суммы для СНИЛС, меньше 1001998, не производится",
			args: args{
				normalizedSnils: "00100199799",
			},
			wantErr: false,
		},
		{
			name: "валидация контрольной суммы для СНИЛС, равный 1001998, не производится",
			args: args{
				normalizedSnils: "00100199899",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.normalizedSnils); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
