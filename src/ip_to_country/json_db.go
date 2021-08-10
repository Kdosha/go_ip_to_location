package ip_to_country

import (
	"encoding/json"
)

type JsonDb struct {
	ipToLocationMap map[string]map[string]string
}

func (db *JsonDb) InitDb(options map[string]interface{}) {
	ipToLocationsJson := `{
		"1.1.1.1": {"country": "asdasd", "city":"a123311"},
		"2.2.2.2": {"country": "asdasd", "city":"a123311"},
		"3.3.3.3": {"country": "asdasd", "city":"a123311"},
		"4.4.4.4": {"country": "asdasd", "city":"a123311"}
    }`

	// Declared an empty map interface
	var result map[string]map[string]string

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(ipToLocationsJson), &result)
	db.ipToLocationMap = result
}

func (db *JsonDb) IpToCountry(ip string) map[string]string {
	return db.ipToLocationMap[ip]
}
