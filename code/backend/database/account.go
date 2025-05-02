package database

import (
	"code/hash_app"
	"code/token_app"
	"code/utils"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateAccount(full_name, contact, email, password, plan string) error {
	full_name = utils.ToLower(utils.Trim(full_name))
	contact = utils.ToLower(utils.Trim(contact))
	email = utils.ToLower(utils.Trim(email))

	err_analysis_validation := ValidationForCreateAccountAnalysis(full_name, contact, email, password)

	if err_analysis_validation != nil {
		return err_analysis_validation
	}

	id := token_app.GenerateUUIDFromEmail(email).String()
	password_hashed, err_password_hashed := hash_app.HashPassword(password)

	if err_password_hashed != nil {
		return err_password_hashed
	}

	db, err_db := GetMongoDB()

	if err_db != nil {
		return err_db
	}

	_, err_find := FindOne(db, LOGIN_STRING, bson.M{"email": email})

	if err_find == nil {
		return fmt.Errorf("conta já existe")
	}

	new_document_login := GenerateDocumentLogin(id, email, password_hashed)

	_, err_insert_login := InsertOne(db, LOGIN_STRING, new_document_login)

	if err_insert_login != nil {
		return err_insert_login
	}

	new_document_manager := GenerateDocumentManager(id, full_name, contact, email, password_hashed, plan)
	_, err_insert_manager_document := InsertOne(db, MANAGER_STRING, new_document_manager)

	if err_insert_manager_document != nil {
		return err_insert_manager_document
	}

	return nil

}

func GenerateDocumentLogin(id, email, password_hashed string) map[string]any {
	datetime, _ := DateNowLocation("", "")

	return map[string]any{
		"id":       strings.ToLower(id),
		"email":    strings.ToLower(email),
		"password": password_hashed,
		"datetime": datetime,
	}

}

func GenerateDocumentManager(id, full_name, contact, email, password_hashed, plan string) map[string]any {
	datetime, _ := DateNowLocation("", "")
	plan = strings.ToLower(plan)

	if plan != "" && !ExistsInList(PLANS, plan) {
		plan = ""
	}

	return map[string]any{
		"id":          strings.ToLower(id),
		"active":      true,
		"plan":        plan,
		"email":       strings.ToLower(email),
		"username":    "",
		"name":        strings.ToLower(full_name),
		"contact":     contact,
		id: map[string]any{
			
		},
		"datetime": datetime,
	}

}

func UpdateUsername(id, username string) error {
	/*
		type attendant_update_value:

		map[string]any{
			"max_lead": <number>,
			"name": <string>,
			"contact": <string>,
		}

	*/

	username = strings.TrimSpace(strings.ToLower(username))

	err_username := ValUsername(username)

	if err_username != nil {
		return err_username
	}

	db, err_db := GetMongoDB()

	if err_db != nil {
		return err_db
	}

	id = strings.TrimSpace(strings.ToLower(id))

	find_username, err_find_username := FindOne(db, MANAGER_STRING, bson.M{USERNAME_STRING: username})

	if find_username != nil || err_find_username == nil {
		return fmt.Errorf("username já existe")
	}

	filter := bson.M{id: bson.M{"$exists": true}}

	update := bson.M{
		"$set": bson.M{
			USERNAME_STRING: username,
		},
	}

	err_update := UpdateOne(db, MANAGER_STRING, filter, update)
	if err_update != nil {
		return err_update
	}

	return nil
}
