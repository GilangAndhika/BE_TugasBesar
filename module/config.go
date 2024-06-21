package module

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoStrinng string =os.Getenv("MONGO_STRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoStrinng,
	DBName: "dbparfume",
}

var MongoConn = atdb.MongoConnect(MongoInfo)