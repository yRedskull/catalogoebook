package database

import (
	/* "code/hash_app" */
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"os"
	"strings"
	"sync"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

var (
	MONGODB_URI               string
	ATTENDANTS_STRING         string = "attendants"
	LEADS_STRING              string = "leads"
	PIXEL_STRING              string = "pixel"
	USERNAME_STRING           string = "username"
	CONTEXT_STRING            string = "context"
	MANAGER_STRING            string = "manager"
	ERROR_STRING              string = "error"
	LOGIN_STRING              string = "login"
	mu                        sync.Mutex
	PLANS                     []string = []string{"basic", "elite"}
	/* PLAN_LIMITATION           bson.M   = bson.M{

		PLANS[0]: bson.M{
			ATTENDANTS_STRING: 4,
			CONTEXT_STRING:    10,
		},
		PLANS[1]: bson.M{
			ATTENDANTS_STRING: 100,
			CONTEXT_STRING:    250,
		},
	} */

	/* HASH_PLANS bson.M = bson.M{
		hash_app.HashString(PLANS[0]): PLANS[0],
		hash_app.HashString(PLANS[1]): PLANS[1],
	} */

	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
	mongoDBName       string
)

func GetMongoDB() (*MongoDB, error) {
	mongoOnce.Do(func() {
		_ = godotenv.Load()

		uri := os.Getenv("MONGODB_URI")
		mongoDBName = os.Getenv("MONGODB_DB_NAME")
		if uri == "" || mongoDBName == "" {
			clientInstanceErr = fmt.Errorf("variáveis de ambiente MONGODB_URI ou MONGODB_DB_NAME não definidas")
			return
		}

		clientOptions := options.Client().ApplyURI(uri)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		clientInstance, clientInstanceErr = mongo.Connect(ctx, clientOptions)
		if clientInstanceErr != nil {
			return
		}

		clientInstanceErr = clientInstance.Ping(ctx, nil)
	})

	if clientInstanceErr != nil {
		return nil, clientInstanceErr
	}

	db := clientInstance.Database(mongoDBName)
	return &MongoDB{Client: clientInstance, Database: db}, nil
}

func InsertOne(db *MongoDB, colletion string, new_document bson.M) (*mongo.InsertOneResult, error) {

	result, err_insert_document := db.Database.Collection(colletion).InsertOne(context.TODO(), new_document)

	if err_insert_document != nil {
		return nil, err_insert_document
	}

	return result, nil
}

func FindOne(db *MongoDB, colletion string, filter bson.M) (bson.M, error) {

	result := db.Database.Collection(colletion).FindOne(context.TODO(), filter)

	if err := result.Err(); err != nil {
		return nil, err
	}

	var value bson.M

	if err := result.Decode(&value); err != nil {
		return nil, err
	}

	return value, nil
}

func Find(ctx context.Context, db *MongoDB, collection string, filter bson.M, fields []string, findOptions ...*options.FindOptions) ([]bson.M, error) {
	col := db.Database.Collection(collection)

	if len(fields) > 0 {
		projection := bson.M{}
		for _, field := range fields {
			projection[field] = 1
		}

		if len(findOptions) > 0 {
			findOptions[0].SetProjection(projection)
		} else {
			findOptions = append(findOptions, options.Find().SetProjection(projection))
		}
	}

	cursor, err := col.Find(ctx, filter, findOptions...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Converte os resultados para um slice de bson.M
	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("nenhum dado encontrado")
	}

	return results, nil
}

func FindAggregate(ctx context.Context, db *MongoDB, collection string, filter bson.M, projection bson.M) ([]bson.M, error) {
	col := db.Database.Collection(collection)

	// Monta o pipeline de agregação: primeiro filtra (match) e depois projeta (project)
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},
		{{Key: "$project", Value: projection}},
	}

	cursor, err := col.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("nenhum dado encontrado")
	}

	return results, nil
}

func UpdateOne(db *MongoDB, collection string, filter bson.M, update bson.M) error {
	_, err := db.Database.Collection(collection).UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}
	return nil
}

func FindByOneEmail(email string) (bson.M, error) {
	db, _err := GetMongoDB()

	if _err != nil {
		return nil, _err
	}

	data_of_user, err := FindOne(db, "login", bson.M{"email": strings.ToLower(email)})

	if err != nil {
		return nil, err
	}

	return data_of_user, nil
}

func FindByOneManager(key, value string) (bson.M, error) {
	db, _err := GetMongoDB()

	if _err != nil {
		return nil, _err
	}

	data_of_manager, err := FindOne(db, MANAGER_STRING, bson.M{key: value})

	if err != nil {
		return nil, err
	}

	return data_of_manager, nil
}

func UpdateManager(id string, key string, newValue any) error {
	db, err := GetMongoDB()

	if err != nil {
		return err
	}

	filter := bson.M{id: bson.M{"$exists": true}}

	update := bson.M{
		"$set": bson.M{
			key: newValue,
		},
	}

	err = UpdateOne(db, MANAGER_STRING, filter, update)
	if err != nil {
		return err
	}

	return nil
}
