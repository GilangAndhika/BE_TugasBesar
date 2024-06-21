package module

import (
	"context"
	// "errors"
	"fmt"
	// "github.com/gryzlegrizz/BE_TugasBesar/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertParfume(db *mongo.Database, col string, nama string, jenis string, merk string, deskripsi string, harga int, thn int, stok int, ukuran string) (insertedID primitive.ObjectID, err error) {
	parfume := bson.M{
		"nama_parfume":    	nama,
		"jenis_parfume":    jenis,
		"merk":     		merk,
		"deskripsi": 		deskripsi,
		"harga":     		harga,
		"thn_peluncuran":   thn,
		"stok":      		stok,
		"ukuran":      		ukuran,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), parfume)
	if err != nil {
		fmt.Printf("InsertParfume: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}