package database

import (
	"context"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func SetPixel(id, pixel_id, access_token string) error {

	id = strings.TrimSpace(strings.ToLower(id))

	db, err_db := GetMongoDB()

	if err_db != nil {
		return err_db
	}

	update_pixel_key := fmt.Sprintf("%s.%s", id, PIXEL_STRING)

	filter := bson.M{id: bson.M{"$exists": true}}

	update := bson.M{
		"$set": bson.M{
			update_pixel_key: map[string]any{
				"id":           pixel_id,
				"access_token": access_token,
			},
		},
	}

	err_update := UpdateOne(db, MANAGER_STRING, filter, update)
	if err_update != nil {
		return err_update
	}

	return nil
}

func ReadPixel(id string) (bson.M, error) {
	db, err_db := GetMongoDB()

	if err_db != nil {
		return nil, err_db
	}

	id = strings.TrimSpace(strings.ToLower(id))

	pixel_key := fmt.Sprintf("%s.%s", id, PIXEL_STRING)

	result, err_find := Find(context.TODO(), db, MANAGER_STRING, bson.M{pixel_key: bson.M{"$exists": true}}, []string{pixel_key})

	if err_find != nil {
		return nil, err_find
	}

	pixel := result[0][id].(bson.M)[PIXEL_STRING]

	if pixel == nil {
		return nil, fmt.Errorf("pixel não existe")

	}

	return pixel.(bson.M), nil
}
