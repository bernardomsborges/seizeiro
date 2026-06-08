package auth

import (
	"errors"
	"testing"
)

func TestValidateCPF(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		cpf   string
		fails bool
	}{
		{
			name: "cpf válido",
			cpf:  "12345678909",
		},
		{
			name: "cpf válido (outro)",
			cpf:  "52998831028",
		},
		{
			name: "cpf válido (com zero à esquerda)",
			cpf:  "01234567890",
		},
		{
			name:  "cpf vazio",
			cpf:   "",
			fails: true,
		},
		{
			name:  "cpf muito curto",
			cpf:   "123456789",
			fails: true,
		},
		{
			name:  "cpf muito longo",
			cpf:   "123456789090",
			fails: true,
		},
		{
			name:  "cpf não normalizado (com pontuação)",
			cpf:   "123.456.789-09",
			fails: true,
		},
		{
			name:  "cpf com caractere não numérico",
			cpf:   "1234567890a",
			fails: true,
		},
		{
			name:  "cpf com dígitos verificadores inválidos",
			cpf:   "12345678900",
			fails: true,
		},
		{
			name:  "cpf com todos os dígitos iguais",
			cpf:   "11111111111",
			fails: true,
		},
		{
			name:  "cpf com todos os dígitos zero",
			cpf:   "00000000000",
			fails: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCPF(tt.cpf)

			if !tt.fails && err != nil {
				t.Fatalf("expected no error for %q, got: %v", tt.cpf, err)
			}
			if tt.fails && err == nil {
				t.Fatalf("expected error for %q", tt.cpf)
			}

			if tt.fails && !errors.Is(err, ErrInvalidCPF) {
				t.Fatalf("expected ErrInvalidCPF for %q, got: %v", tt.cpf, err)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		senha          string
		fails          bool
		wantViolations []string
	}{
		{
			name:  "senha forte",
			senha: "Abc12345",
		},
		{
			name:  "senha forte (com caractere especial)",
			senha: "SenhaForte_2026",
		},
		{
			name:           "senha muito curta",
			senha:          "Ab1",
			fails:          true,
			wantViolations: []string{"entre 8 e 50 caracteres"},
		},
		{
			name:           "senha fraca (sem maiúscula e dígito)",
			senha:          "soletraminuscula",
			fails:          true,
			wantViolations: []string{"uma letra maiúscula", "um dígito"},
		},
		{
			name:           "senha fraca (sem minúscula)",
			senha:          "SOMAISCULA123",
			fails:          true,
			wantViolations: []string{"uma letra minúscula"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.senha)

			if !tt.fails && err != nil {
				t.Fatalf("expected not error for %q, got: %v", tt.senha, err)
			}
			if tt.fails && err == nil {
				t.Fatalf("expected error for %q", tt.senha)
			}

			if tt.fails {
				weakErr, ok := errors.AsType[*WeakPasswordError](err)
				if !ok {
					t.Fatalf("expected *WeakPasswordError, got %v", err)
				}

				if len(weakErr.Violations) != len(tt.wantViolations) {
					t.Errorf("invalid lenght. want %d, got %d", len(tt.wantViolations), len(weakErr.Violations))
					return
				}
			}
		})
	}
}
