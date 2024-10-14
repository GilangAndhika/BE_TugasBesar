package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/GilangAndhika/BE_TugasBesar/model"

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
		"nama_parfume":    		nama,
		"jenis_parfume":    	jenis,
		"merk":     			merk,
		"deskripsi": 			deskripsi,
		"harga":     			harga,
		"tahun_peluncuran":   	thn,
		"stok":      			stok,
		"ukuran":      			ukuran,
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
			"nama_parfume":    		nama,
			"jenis_parfume":    	jenis,
			"merk":     			merk,
			"deskripsi": 			deskripsi,
			"harga":     			harga,
			"tahun_peluncuran":   	thn,
			"stok":      			stok,
			"ukuran":      			ukuran,
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

func InsertUser(db *mongo.Database, col string, username string, password string, idrole primitive.ObjectID, email string, phone string, address string) (insertedID primitive.ObjectID, err error) {
	user := bson.M{
		"username": username,
		"password": password,
		"idrole":  primitive.NewObjectID(),
		"email":    email,
		"phone":    phone,
		"address":  address,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("InsertUser: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database, col string) (user model.User, errs error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, fmt.Errorf("no data found for ID %s", _id)
		}
		return user, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return user, nil
}

func UpdateUser(_id primitive.ObjectID, db *mongo.Database, col string, username string, password string, idrole primitive.ObjectID, email string, phone string, address string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"username": username,
			"password": password,
			"idrole":  idrole,
			"email":    email,
			"phone":    phone,
			"address":  address,
		},
	}
	result, err :=db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateUser: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data updated with the specified ID")
		return
	}
	return nil
}

func DeleteUserByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
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

func GetAllUser(db *mongo.Database, col string) (data []model.User) {
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
func InsertRole(db *mongo.Database, col string, roleUser string) (insertedID primitive.ObjectID, err error) {
	role := bson.M{
		"role": roleUser,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), role)
	if err != nil {
		fmt.Printf("InsertRole: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetRoleFromID(_id primitive.ObjectID, db *mongo.Database, col string) (role model.Roles, errs error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := collection.FindOne(context.TODO(), filter).Decode(&role)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return role, fmt.Errorf("no data found for ID %s", _id)
		}
		return role, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return role, nil
}

func UpdateRole(_id primitive.ObjectID, db *mongo.Database, col string, roleUser string) (err error) {
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": bson.M{
			"role": roleUser,
		},
	}
	result, err :=db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateRole: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data updated with the specified ID")
		return
	}
	return nil
}

func DeleteRoleByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
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

func GetAllRole(db *mongo.Database, col string) (data []model.Roles) {
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

// func BuatToken(db *mongo.Database, col string, token string, username string) (insertedID primitive.ObjectID, err error) {
// 	user := bson.M{
// 		"token": token,
// 		"username": username,
// 	}
// 	result, err := db.Collection(col).InsertOne(context.Background(), user)
// 	if err != nil {
// 		fmt.Printf("BuatToken: %v\n", err)
// 		return
// 	}
// 	insertedID = result.InsertedID.(primitive.ObjectID)
// 	return insertedID, nil
// }