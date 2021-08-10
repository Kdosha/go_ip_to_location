package ip_to_country

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type csvDb struct {
	ipToLocationMap map[string]map[string]string
}

func (db *csvDb) InitDb(options map[string]interface{}) {
	filePath := options["file_path"].(string)
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvFile)
	result := make(map[string]map[string]string)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		ip, city, country := record[0], record[1], record[2]
		result[ip] = map[string]string{"city": city, "country": country}
	}
	db.ipToLocationMap = result
}

func (db *csvDb) IpToCountry(ip string) map[string]string {
	return db.ipToLocationMap[ip]
}
