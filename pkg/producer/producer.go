package producer

import (
	"context"
	"msisdn/pkg/entities"
	"msisdn/pkg/pb"
	"time"
)

type MsisdnProducer interface {
	GetMsisdnDetails(ctx context.Context, input string) (*entities.Msisdn, error)
}

type msisdnProducer struct {
	msisdnSvc pb.MsisdnServiceClient
}

func NewMsisdnProducer(msisdnSvc pb.MsisdnServiceClient) MsisdnProducer {
	return &msisdnProducer{
		msisdnSvc: msisdnSvc,
	}
}

func (m *msisdnProducer) GetMsisdnDetails(ctx context.Context, input string) (*entities.Msisdn, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := m.msisdnSvc.GetMsisdnDetails(ctx, &pb.MsisdnRequest{Msisdn: input})
	if err != nil {
		return nil, err
	}

	return &entities.Msisdn{
		Mno:         resp.GetMno(),
		Cdc:         resp.GetCdc(),
		Sn:          resp.GetSn(),
		CountryCode: resp.GetCountryCode(),
		CountryName: resp.GetCountryName(),
	}, nil

}
