package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/gryzlegrizz/BE_TugasBesar/model"

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

func GetParfumeFromID(_id primitive.ObjectID, db *mongo.Database, col string) (parfume model.Parfume, errs error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&parfume)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return parfume, fmt.Errorf("no data found for ID %s", _id)
		}
		return parfume, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return parfume, nil
}

func GetAllParfume(db *mongo.Database, col string) (data []model.Parfume) {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func UpdateParfume(_id primitive.ObjectID, db *mongo.Database, col string, nama string, jenis string, merk string, deskripsi string, harga int, thn int, stok int, ukuran string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"nama_parfume":    	nama,
			"jenis_parfume":    jenis,
			"merk":     		merk,
			"deskripsi": 		deskripsi,
			"harga":     		harga,
			"thn_peluncuran":   thn,
			"stok":      		stok,
			"ukuran":      		ukuran,
		},
	}
	result, err :=db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateParfume: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data updated with the specified ID")
		return
	}
	return nil
}

func DeleteParfumeByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil{
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}
	if result.DeletedCount == 0	{
		return fmt.Errorf("data with ID %s not found", _id)
	}
	return nil
}