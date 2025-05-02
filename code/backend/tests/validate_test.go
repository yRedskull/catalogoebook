package tests

import (
	"code/database"
	"code/hash_app"
	"testing"
)

func TestValEmail(t *testing.T) {
	email := "tainan.bussiness@gmail.com"
	result := database.ValEmail(email)

	if result != nil {
		t.Errorf("O %s está inválido.", email)
	}
}

func TestValCPF(t *testing.T) {
	cpf := "02815787962"
	result := database.ValidateCPF(cpf)

	if result != nil {
		t.Errorf("O %s está inválido.", cpf)
	}
}

func TestAnalyzeValidationForCreateAccount(t *testing.T) {
	full_name := "Tainan Felipe"
	email := "tainan.bussiness@gmail.com"
	contact := "44988604098"
	password := "12634"

	result := database.ValidationForCreateAccountAnalysis(full_name, contact, email, password)

	if result != nil {
		t.Errorf("AnalyzeValidationForCreateAccount retornou erro: %s", result)
	}
}

func TestValUsername(t *testing.T) {
	username := "tainan"

	result := database.ValUsername(username)

	if result != nil {
		t.Errorf("ValUsername retornou erro: %s", result)
	}
}

func TestHashString(t *testing.T) {
	str := "2"

	result := hash_app.HashString(str)

	if result == "" {
		t.Errorf("HashString retornou nada")
	}
}