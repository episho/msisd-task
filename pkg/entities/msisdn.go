package entities

type Msisdn struct {
	Mno         string //mobile network operator identifier
	Cdc         string //country dialling code
	Sn          string //subscriber number
	CountryCode string //country identifier ISO 3166-1-alpha-2
	CountryName string
}

