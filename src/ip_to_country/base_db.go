// Package ip_to_country - every new DB should use this interface
package ip_to_country

type IpToLocationDb interface {
	InitDb(options map[string]interface{})
	IpToCountry(ip string) map[string]string
}
