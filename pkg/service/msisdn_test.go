package service

import (
	"context"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
	"time"

	"msisdn/pkg/entities"
	"msisdn/pkg/pb"
)

type MsisdnServerTestSuite struct {
	suite.Suite
	msisdnSvcServer pb.MsisdnServiceServer
	testData        []struct {
		input       string
		expected    *entities.Msisdn
		expectedErr error
	}
}

func (suite *MsisdnServerTestSuite) SetupSuite() {
	suite.msisdnSvcServer = NewMsisdnServer()
	os.Setenv("DATA_PATH", "../data")

	// expected response
	suite.testData = []struct {
		input       string
		expected    *entities.Msisdn
		expectedErr error
	}{
		{
			"00389 71 233 444",
			&entities.Msisdn{
				Mno:         "Makedonski Telekom",
				Cdc:         "389",
				Sn:          "71233444",
				CountryCode: "MK",
				CountryName: "Macedonia",
			},
			nil,
		},
		{
			"+(389) 75-789-678",
			&entities.Msisdn{
				Mno:         "A1 Macedonia",
				Cdc:         "389",
				Sn:          "75789678",
				CountryCode: "MK",
				CountryName: "Macedonia",
			},
			nil,
		},
		{
			"+(111) 75-789-678",
			nil,
			entities.ErrInvalidDialingCode,
		},
		{
			"38971232954867596895064065064-56-3597568",
			nil,
			entities.ErrInvalidLenghtMsisdnInput,
		},
		{
			"+(389) 75-789-6781",
			nil,
			entities.ErrInvalidLengthNumber,
		},
		{
			"00389 73-111-222",
			nil,
			entities.ErrInvalidMnoCode,
		},
		{
			"00389 073-111-222",
			nil,
			entities.ErrInvalidLengthNumber,
		},
	}
}

func (suite *MsisdnServerTestSuite) TestGetMsisdnDetails() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, test := range suite.testData {
		resp, err := suite.msisdnSvcServer.GetMsisdnDetails(ctx, &pb.MsisdnRequest{
			Msisdn: test.input,
		})

		if err != nil {
			suite.EqualError(err, test.expectedErr.Error())

			continue
		}

		suite.EqualMsisdn(test.expected, resp)
	}

}

func (suite *MsisdnServerTestSuite) EqualMsisdn(exp *entities.Msisdn, resp *pb.Msisdn) {
	suite.Equal(exp.CountryName, resp.GetCountryName())
	suite.Equal(exp.CountryCode, resp.GetCountryCode())
	suite.Equal(exp.Sn, resp.GetSn())
	suite.Equal(exp.Cdc, resp.GetCdc())
	suite.Equal(exp.Mno, resp.GetMno())
}

func TestIntegrationMsisdnServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsisdnServerTestSuite))
}
