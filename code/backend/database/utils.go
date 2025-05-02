package database

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	VAL_ITEMS = map[string]any{
		"max_lead": []any{ValMaxLead, "int32"},
		"name":     []any{ValName, "string"},
		"contact":  []any{ValContact, "string"},
		"username": []any{ValUsername, "string"},
		"email":    []any{ValEmail, "string"},
	}
	LIST_ITEM_ACCEPTABLE_ATTENDANT = []string{"max_lead", "name", "contact"}
	LIST_ITEM_ACCEPTABLE_MESSAGE   = []string{"text-message"}
)

func ConvertInt32(number any) int32 {
	switch v := number.(type) {
	case int:
		return int32(v)
	case int64:
		return int32(v)
	case int32:
		return v
	default:
		return int32(0)
	}
}

func ValidationItem(validation_item []any, new_item_value any) (any, error) {
	var (
		err_item   error
		item_value any = new_item_value
	)

	switch validation_item[1] {
	case "string":
		if !ValidateType(new_item_value, "string") {
			return nil, fmt.Errorf("tipo do valor inválido")
		}

		item_value = strings.TrimSpace(strings.ToLower(new_item_value.(string)))
		err_item = validation_item[0].(func(string) error)(new_item_value.(string))

	case "int32":
		if !ExistsInList([]string{"int32", "int64", "int"}, reflect.TypeOf(new_item_value).String()) {
			return nil, fmt.Errorf("tipo do valor inválido")
		}

		item_value = ConvertInt32(new_item_value)

		err_item = validation_item[0].(func(int32) error)(item_value.(int32))
	}

	if err_item != nil {
		return nil, err_item
	}

	return item_value, nil
}

func ValidateType(first, second any) bool {
	first_type := reflect.TypeOf(first)
	second_type := reflect.TypeOf(second)

	return first_type == second_type
}

func ValUsername(username string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9]{1,50}$`)
	validation := re.MatchString(username)

	if !validation {
		return fmt.Errorf("identificador inválido")
	}

	return nil
}

func ValMaxLead(value int32) error {
	if value > 10000 || value < 1 {
		return fmt.Errorf("insira um número inteiro entre 1 e 10000")
	}

	return nil
}

func ValName(value string) error {
	re := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s]+$`)
	validation := re.MatchString(value)

	if !validation {
		return fmt.Errorf("nome inválido")
	}

	return nil
}

func ValFullName(value string) error {
	err_name := ValName(value)

	if err_name != nil {
		return err_name
	}

	if len(strings.Split(value, " ")) < 2 {
		return fmt.Errorf("insira o nome completo")
	}

	return nil
}

func ValContact(value string) error {
	re := regexp.MustCompile(`^\d{2}9?\d{8}$`)
	validation := re.MatchString(value)

	if !validation {
		return fmt.Errorf("número inválido")
	}

	return nil
}

func ValEmail(value string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	validation := re.MatchString(value)

	if !validation {
		return fmt.Errorf("email inválido")
	}

	return nil
}

func ValPassword(value string) error {
	if len(value) < 5 {
		return fmt.Errorf("senha deverá conter 5 ou mais caracteres")
	}

	return nil
}

func ValidateCPF(cpf string) error {

	type_error := fmt.Errorf("erro de tipagem")

	re := regexp.MustCompile(`\D`)
	strCPF := re.ReplaceAllString(cpf, "")

	// O CPF deve ter exatamente 11 dígitos.
	if len(strCPF) != 11 {
		return fmt.Errorf("não contém 11 dígitos")
	}

	// Verifica se o CPF é uma sequência de números repetidos.
	repetidos := []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
	}
	for _, rep := range repetidos {
		if strCPF == rep {
			return fmt.Errorf("são valores repetidos")
		}
	}

	sum := 0
	for i := 1; i <= 9; i++ {
		digit, err := strconv.Atoi(strCPF[i-1 : i])
		if err != nil {
			return type_error
		}
		sum += digit * (11 - i)
	}

	rest := (sum * 10) % 11
	if rest == 10 || rest == 11 {
		rest = 0
	}

	firstDigit, err := strconv.Atoi(strCPF[9:10])
	if err != nil || rest != firstDigit {
		return fmt.Errorf("cpf inválido")
	}

	sum = 0
	for i := 1; i <= 10; i++ {
		digit, err := strconv.Atoi(strCPF[i-1 : i])
		if err != nil {
			return type_error
		}
		sum += digit * (12 - i)
	}

	rest = (sum * 10) % 11
	if rest == 10 || rest == 11 {
		rest = 0
	}

	secondDigit, err := strconv.Atoi(strCPF[10:11])
	if err != nil || rest != secondDigit {
		return fmt.Errorf("cpf inválido")
	}

	return nil
}

func ExistsInList(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func ValidationForCreateAttendantAnalysis(id, attendant_id, attendant_name, contact_number string, max_lead any) error {
	if id == "" {
		return fmt.Errorf("id do gerente inválido")
	}

	err_attendant_id := ValUsername(attendant_id)

	if attendant_id == "" || err_attendant_id != nil {
		return fmt.Errorf("id do atendente inválido")
	}

	err_attendant_name := ValName(attendant_name)

	if attendant_name == "" || err_attendant_name != nil {
		return fmt.Errorf("nome do atendente inválido")
	}

	err_contact_number := ValContact(contact_number)

	if contact_number == "" || err_contact_number != nil {
		return fmt.Errorf("contato do atendente inválido")
	}

	max_leadint32 := ConvertInt32(max_lead)

	err_max_lead := ValMaxLead(max_leadint32)

	if max_lead == 0 || err_max_lead != nil {
		return fmt.Errorf("número de redirecionamento do atendente inválido")
	}

	return nil
}

func ValidationForCreateLeadAnalysis(id, lead_name, lead_email, lead_contact string) error {
	if id == "" {
		return fmt.Errorf("id do gerente inválido")
	}

	exists_name := lead_name == ""
	exists_email := lead_email == ""
	exists_contact := lead_contact == ""

	if !(exists_name) {
		err_lead_name := ValName(lead_name)

		if err_lead_name != nil {
			return fmt.Errorf("nome inválido")
		}
	}

	if !(exists_email) {
		err_lead_email := ValEmail(lead_email)

		if err_lead_email != nil {
			return fmt.Errorf("email inválido")
		}
	}

	if !(lead_contact == "") {

		err_contact_number := ValContact(lead_contact)

		if err_contact_number != nil {
			return fmt.Errorf("contato inválido")
		}
	}

	if exists_contact && exists_email {
		return fmt.Errorf("lead inválido")
	}

	return nil
}
func ValidationForContextMessageAnalysis(id, id_message, message string) error {
	if id == "" {
		return fmt.Errorf("id do gerente inválido")
	}

	err_id_message := ValUsername(id_message)

	if id_message == "" || err_id_message != nil {
		return fmt.Errorf("id da mensagem inválido")
	}

	if message == "" {
		return fmt.Errorf("mensagem inválida")
	}

	return nil
}

func ValidationForCreateAccountAnalysis(full_name, contact, email, password string) error {

	err_full_name := ValFullName(full_name)

	if full_name == "" || err_full_name != nil {
		return err_full_name
	}

	err_contact_number := ValContact(contact)

	if contact == "" || err_contact_number != nil {
		return err_contact_number
	}

	err_email := ValEmail(email)

	if email == "" || err_email != nil {
		return err_email
	}

	err_password := ValPassword(password)

	if password == "" || err_password != nil {
		return err_password
	}

	return nil
}

func DateNowLocation(location, format string) (string, error) {
	if location == "" {
		location = "America/Sao_Paulo"
	}

	if format == "" {
		format = "02/01/2006 15:04:05"
	}

	loc, err_location := time.LoadLocation(location)

	if err_location != nil {
		return "", err_location
	}

	now := time.Now().In(loc)

	return now.Format(format), nil
}

func PlanRestrictionAnalysis(plan, field string, manager any) {

}

func FindObjInList(obj any, list []any) *any {
	for i, p := range list {
		if p == obj {
			return &list[i]
		}
	}
	return nil
}
