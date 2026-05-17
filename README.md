# snils

[![Go Version](https://img.shields.io/badge/go-1.26.2-00ADD8.svg)](https://go.dev)
[![CI](https://github.com/Kenedy228/snils/actions/workflows/ci.yml/badge.svg)](https://github.com/Kenedy228/snils/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/Kenedy228/snils)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/Kenedy228/snils.svg)](https://pkg.go.dev/github.com/Kenedy228/snils)

Go-библиотека для нормализации и валидации СНИЛС (страхового номера индивидуального лицевого счета) Российской Федерации.

## ✨ Возможности

- нормализация СНИЛС;
- проверка длины;
- проверка содержимого;
- проверка запрещенных значений;
- проверка контрольной суммы по официальному алгоритму;
- поддержка `errors.Is`.

## 📦 Установка

```bash
go get github.com/Kenedy228/snils
```

## 🚀 Использование

```go
package main

import (
	"errors"
	"fmt"

	"github.com/Kenedy228/snils"
)

func main() {
	raw := "112-233-445 95"

	normalized := snils.Normalize(raw)

	if err := snils.Validate(normalized); err != nil {
		fmt.Println("СНИЛС невалиден:", err)

		if errors.Is(err, snils.ErrInvalidChecksum) {
			fmt.Println("Неверная контрольная сумма")
		}

		return
	}

	fmt.Println("СНИЛС валиден")
}
```

---

## 🧩 API

### Нормализация

```go
normalized := snils.Normalize("112-233-445 95")

fmt.Println(normalized)
```

Результат:

```text
11223344595
```

### Валидация

```go
err := snils.Validate("11223344595")
fmt.Println(err == nil)
```

Результат:

```text
true
```

### Ошибки

Пакет возвращает следующие ошибки:

| Ошибка | Описание |
|---|---|
| `snils.ErrInvalidLength` | Неверная длина СНИЛС |
| `snils.ErrInvalidContent` | СНИЛС содержит недопустимые символы |
| `snils.ErrForbiddenSNILS` | Запрещенное значение СНИЛС |
| `snils.ErrInvalidChecksum` | Неверная контрольная сумма |

Пример обработки:

```go
if errors.Is(err, snils.ErrInvalidChecksum) {
	// обработка ошибки контрольной суммы
}
```

### Формат СНИЛС

Нормализованный СНИЛС состоит из 11 цифр:

```text
XXX XXX XXX YY
```

где:

- `XXX XXX XXX` — номер
- `YY` — контрольное число

### Контрольная сумма

Контрольное число рассчитывается следующим образом:

1. Каждая из первых 9 цифр умножается на вес от 9 до 1;
2. Полученные значения суммируются;
3. Применяются правила расчета контрольного числа:
    - если сумма меньше 100 — контрольное число равно сумме;
    - если сумма равна 100 — контрольное число равно 0;
    - если сумма больше 100:
        - вычисляется остаток от деления на 101;
        - если остаток равен 100 — контрольное число равно 0;
        - иначе контрольное число равно остатку.

Проверка контрольного числа выполняется только для номеров больше `1001998`.

## 🧪 Тестирование

```bash
go test ./...
```

---

## 🔗 Источники

- https://www.consultant.ru/document/cons_doc_LAW_142584/1d9155a863a5949b14b95ecbb536aa84856a2a2e/
- https://info.gosuslugi.ru/articles/Валидация/

---

## 📄 Лицензия

MIT