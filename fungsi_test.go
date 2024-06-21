package _tes

import (
	"fmt"
	"testing"

	// "github.com/gryzlegrizz/BE_TugasBesar/model"
	"github.com/gryzlegrizz/BE_TugasBesar/module"
	// "go.mongodb.org/mongo-driver/bson/primitive"
) 

func TestInsertParfume(t* testing.T){
	nama := 		"Chirstian Dior"
	jenis := 		"Eau de Parfum"
	merk := 		"Dior"
	deskripsi := 	"Parfum yang sangat wangi"
	harga :=		1000000
	thn := 			2021
	stok := 		100
	ukuran := 		"100ml"
	insertedID, err := module.InsertParfume(module.MongoConn, "parfume", nama, jenis, merk, deskripsi, harga, thn, stok, ukuran)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan ID: %v\n", insertedID.Hex())
}

