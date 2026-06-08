package auth

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	cpfReplacer = strings.NewReplacer(".", "", "-", "")
)

// Normaliza um nome removendo espaços em excesso.
func normalizeNome(nome string) string {
	return strings.Join(strings.Fields(nome), " ")
}

// Normalize um CPF removendo pontos e hífens.
func normalizeCPF(cpf string) string {
	return cpfReplacer.Replace(cpf)
}

// Normaliza um email removendo espaços e forçando caixa baixa.
func normalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}

// ValidateCPF valida um CPF. O valor informado deve conter apenas dígitos, sem formatação.
// Verifica o comprimento, descarta sequências de dígitos repetidos e confere os dígitos verificadores.
//
// Retorna [ErrInvalidCPF] caso o CPF seja inválido.
func ValidateCPF(cpf string) error {
	if len(cpf) != 11 {
		return ErrInvalidCPF
	}

	var digits [11]int
	allEqual := true
	for i := range 11 {
		c := cpf[i]
		if c < '0' || c > '9' {
			return ErrInvalidCPF
		}
		digits[i] = int(c - '0')
		if digits[i] != digits[0] {
			allEqual = false
		}
	}

	// Dígitos repetidos passam na validação mas são inválidos.
	if allEqual {
		return ErrInvalidCPF
	}

	for i := 9; i < 11; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += digits[j] * (i + 1 - j)
		}

		check := (sum * 10) % 11
		if check == 10 {
			check = 0
		}

		if check != digits[i] {
			return ErrInvalidCPF
		}
	}

	return nil
}

// ValidatePassword valida a força de uma senha. A senha deve possuir entre 8 a 50 caracteres,
// um dígito, uma letra minúscula e uma maiúscula.
// Retorna [*WeakPasswordError] caso a senha seja fraca.
func ValidatePassword(senha string) error {
	var (
		lower bool
		upper bool
		digit bool
	)

	var violations []string

	n := utf8.RuneCountInString(senha)
	if n < 8 || n > 50 {
		violations = append(violations, "entre 8 e 50 caracteres")
	}

	for _, r := range senha {
		switch {
		case unicode.IsLower(r):
			lower = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsDigit(r):
			digit = true
		}
	}

	if !lower {
		violations = append(violations, "uma letra minúscula")
	}
	if !upper {
		violations = append(violations, "uma letra maiúscula")
	}
	if !digit {
		violations = append(violations, "um dígito")
	}

	if len(violations) > 0 {
		return &WeakPasswordError{
			Violations: violations,
		}
	}

	return nil
}
