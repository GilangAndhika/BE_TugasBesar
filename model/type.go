package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type parfume struct {
	ID				primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama_Parfume	string				`bson:"nama_parfume,omitempty" json:"nama_parfume,omitempty" example:"Chirstian Dior"`
	Jenis_Parfume	string				`bson:"jenis_parfume,omitempty" json:"jenis_parfume,omitempty" example:"Eau de Parfum"`
	Merk			string				`bson:"merk,omitempty" json:"merk,omitempty" example:"Dior"`
	Deskripsi		string				`bson:"deskripsi,omitempty" json:"deskripsi,omitempty" example:"Parfum yang sangat wangi"`
	Harga			int					`bson:"harga,omitempty" json:"harga,omitempty" example:"1000000"`
	Thn_Peluncuran	int					`bson:"tahun_peluncuran,omitempty" json:"tahun_peluncuran,omitempty" example:"2021"`
	Stok			int					`bson:"stok,omitempty" json:"stok,omitempty" example:"100"`
	Ukuran			string				`bson:"ukuran,omitempty" json:"ukuran,omitempty" example:"100ml"`
}

type user struct {
	ID				primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Username		string				`bson:"username,omitempty" json:"username,omitempty" example:"user"`
	Password		string				`bson:"password,omitempty" json:"password,omitempty" example:"user"`
	Email			string				`bson:"email,omitempty" json:"email,omitempty" example:"user"`
	Phone			string				`bson:"phone,omitempty" json:"phone,omitempty" example:"08123456789"`
	Address			string				`bson:"address,omitempty" json:"address,omitempty" example:"Jl. Jalan"`
}