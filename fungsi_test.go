package _tes

import (
	"fmt"
	"testing"

	// "github.com/gryzlegrizz/BE_TugasBesar/model"
	"github.com/gryzlegrizz/BE_TugasBesar/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
) 

func TestInsertParfume(t* testing.T){
	nama := 		"Van Persie"
	jenis := 		"Eau de Parfum"
	merk := 		"Chanel"
	deskripsi := 	"Parfum yang sangat harum dan tahan lama"
	harga :=		1500000
	thn := 			2019
	stok := 		100
	ukuran := 		"50ml"
	insertedID, err := module.InsertParfume(module.MongoConn, "parfume", nama, jenis, merk, deskripsi, harga, thn, stok, ukuran)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan ID: %v\n", insertedID.Hex())
}

func TestGetParfumeFromID(t *testing.T) {
    id := "66756b2bb0584a360d9709d8"
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        t.Fatalf("Error converting id to ObjectID: %v", err)
    }
    parfume, err := module.GetParfumeFromID(objectID, module.MongoConn, "parfume")
    if err != nil {
        t.Fatalf("Error calling GetParfumeFromID: %v", err)
    }
	fmt.Println(parfume)
}

func TestGetAllParfume(t *testing.T) {
	data := module.GetAllParfume(module.MongoConn, "parfume")
	fmt.Println(data)
}