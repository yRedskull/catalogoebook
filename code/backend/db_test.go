package main

import (
	"code/database"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	full_name := "Tainan Felipe Dos Santos"
	contact := "44988604098"
	email := "tainan.bussiness@gmail.com"
	password := "12345"

	err_create := database.CreateAccount(full_name, contact, email, password, "basic")

	if err_create != nil {
		t.Errorf("Função CreateAccount retornou erro: %s", err_create.Error())
	}
}

func TestDateNowLocation(t *testing.T) {
	_, err := database.DateNowLocation("", "")

	if err != nil {
		t.Errorf("Função DateNowLocation não funcionou")
	}
}

/* func TestUpdateUsername(t *testing.T) {
	err := database.UpdateUsername("14935257946", "tainan")

	if err != nil {
		t.Errorf("Função UpdateUsername retornou erro: %s", err.Error())
	}
} */

// CRUD Context Message
/* 
func TestCreateContextMessage(t *testing.T) {
	err := database.CreateContextMessage("14935257946", "alfa", "Olá, quero saber mais sobre o Alfa")

	if err != nil {
		t.Errorf("Função CreateContextMessage retornou erro: %s", err.Error())
	}
} 

func TestReadContextMessage(t *testing.T) {
	result, err := database.ReadContextMessage("14935257946", "alfa")

	if err != nil {
		t.Errorf("Função ReadContextMessage retornou erro: %s", err.Error())
	}

	fmt.Println(result)
}

func TestUpdateContextMessage(t *testing.T) {
	err := database.UpdateContextMessage("14935257946", "bonivita", "Olá, vi o seu anuncio do Bonivita e quero saber mais sobre.")

	if err != nil {
		t.Errorf("Função UpdateContextMessage retornou erro: %s", err.Error())
	}
} 

func TestDeleteContextMessage(t *testing.T) {
	err := database.DeleteContextMessage("14935257946", "bonivita")

	if err != nil {
		t.Errorf("Função DeleteContextMessage retornou erro: %s", err.Error())
	}
}

 */
/* 
 func TestToggleStatusAttendant(t *testing.T) {
	err := database.ToggleStatusAttendant("14935257946", "tainan", false)

	if err != nil {
		t.Errorf("Função ToggleStatusAttendant retornou erro: %s", err.Error())
	}
}
 */

/*  func TestCreateLead(t *testing.T) {
	h_plan := hash_app.HashString("basic")
	err := database.CreateLead("14935257946", h_plan, "Priscila Maciel", "priscila15maciel@gmail.com", "44997071285", "Interesse em vendas", "tainan")

	if err != nil {
		t.Errorf("Função ReadContextMessage retornou erro: %s", err.Error())
	}
} */
 /* func TestReadLead(t *testing.T) {
	_, err := database.ReadLead("14935257946", "contact", "44997071265")

	if err != nil {
		t.Errorf("Função ReadLead retornou erro: %s", err.Error())
	}
} */

/* func TestCreateLogError(t *testing.T){
	_, err_read := database.ReadLead("14234214535325", "conta2ct", "449970715265")

	err := database.CreateLogError("14935257946", err_read)

	if err != nil {
		t.Errorf("Função CreateLogError retornou erro: %s", err.Error())
	}
} */


