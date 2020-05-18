package service

import (
	"context"
	"fmt"

	"strings"

	"msisdn/pkg/entities"
	"msisdn/pkg/pb"
)

type msisdnSvcServer struct{}

func NewMsisdnServer() pb.MsisdnServiceServer {
	return &msisdnSvcServer{}
}
func (m *msisdnSvcServer) GetMsisdnDetails(ctx context.Context, req *pb.MsisdnRequest) (*pb.Msisdn, error) {
	data := new(entities.MsisdnData)

	//err:=os.Setenv("data_path", "../pkg/data")
	//if err!=nil{
	//	fmt.Println("failed to set data_path value")
	//}

	err := data.LoadData()
	if err != nil {
		return nil, fmt.Errorf("failed to get msisdn data error: %v", err)
	}

	cleanInput, err := entities.CleanMsisdnInput(req.Msisdn)
	if err != nil {
		return nil, err
	}

	cc := getCountryCode(data, cleanInput)

	if cc == (entities.CC{}) {
		return nil, entities.ErrInvalidDialingCode
	}

	// remove the dialing code eg.for MKD is 389
	inputTrimDialCode := strings.TrimPrefix(cleanInput, cc.DialCode)

	// check the length of subscriber number
	if len(inputTrimDialCode) != cc.SizeOfNN {
		return nil, entities.ErrInvalidLengthNumber
	}

	operator := getMobileNetworkOperator(data, inputTrimDialCode)
	if len(operator) == 0 {
		return nil, entities.ErrInvalidMnoCode
	}

	return &pb.Msisdn{
		CountryName: cc.Name,
		CountryCode: cc.Code,
		Mno:         operator,
		Cdc:         cc.DialCode,
		Sn:          inputTrimDialCode,
	}, nil
}

func getCountryCode(m *entities.MsisdnData, input string) entities.CC {
	for _, c := range m.CountryData {
		dialCode := input[:len(c.DialCode)]

		if dialCode == c.DialCode {
			return c
		}
	}
	return entities.CC{}
}

func getMobileNetworkOperator(m *entities.MsisdnData, input string) string {
	for _, mno := range m.MnoData {
		for _, code := range mno.Code {
			inputCode := input[:len(code)]

			if inputCode == code {
				return mno.Operator
			}
		}
	}

	return ""
}
