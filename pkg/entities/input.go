package entities

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
)



type MsisdnData struct {
	CountryData []CC
	MnoData     []Mno
}

type CC struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	DialCode string `json:"dial_code"`
	SizeOfNN int    `json:"size_of_nn"` //number of digits that includes the mobile prefix and excludes the country code
}

type Mno struct {
	Operator string   `json:"operator"`
	Code     []string `json:"code"`
}

func (m *MsisdnData) LoadData() error {
	os.Setenv("DATA_PATH","../pkg/data")
	data_path:=os.Getenv("DATA_PATH")
	countryFile := data_path + "/country.json"
	countryJSON, err := handleFile(countryFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(countryJSON, &m.CountryData); err != nil {
		return err
	}

	mnoDataFile := data_path + "/mkd-mno.json"
	mnoJSON, err := handleFile(mnoDataFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(mnoJSON, &m.MnoData); err != nil {
		return err
	}

	return nil
}

// handleFile checks if file exists, open and load it
func handleFile(filepath string) ([]byte, error) {
	_, err := CheckFile(filepath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, err
}

// Check if file exist in directory.
func CheckFile(filepath string) (bool, error) {

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, err
	}

	return true, nil
}

// cleanInput cleans the input
func CleanMsisdnInput(input string) (cleanInput string, err error) {

	/* with this expresion we remove:
		- remove the first one or two zeros
		- remove dashes -
		- remove plus +
		- remove left and right parenthesis ( )
		- remove a whitespace character: [\t\n\f\r]
	*/
	reg := regexp.MustCompile(`^0{1,2}|-|\+|\(|\)|\s`)
	cleanInput = reg.ReplaceAllString(input, "")

	//check if it contains only digits
	reg = regexp.MustCompile("^[0-9]+$")
	ok := reg.MatchString(cleanInput)
	if !ok {
		return "", ErrInputHasLetters
	}

	// The ITU-T recommendation E.164 limits the maximum length of an MSISDN from 8 to 15 digits
	r := regexp.MustCompile("^[0-9]{8,15}$")
	ok = r.MatchString(cleanInput)
	if !ok {
		return "", ErrInvalidLenghtMsisdnInput
	}

	return cleanInput, nil
}

