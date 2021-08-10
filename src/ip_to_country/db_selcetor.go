package ip_to_country

import (
	"learning_go/configs"
	"log"
)

func GetDbObj(options map[string]interface{}) IpToLocationDb {
	var db IpToLocationDb
	switch configs.DbType {
	case "csv":
		db = &csvDb{}
	case "json":
		db = &JsonDb{}
	default:
		log.Fatalln("Invalid DB type")
	}

	db.InitDb(options)
	return db
}
