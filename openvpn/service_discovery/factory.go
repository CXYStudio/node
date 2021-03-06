package service_discovery

import (
	"github.com/mysterium/node/datasize"
	dto "github.com/mysterium/node/openvpn/service_discovery/dto"
	dto_discovery "github.com/mysterium/node/service_discovery/dto"
	"time"
)

func NewServiceProposal(nodeKey string, nodeLocation dto_discovery.Location) dto_discovery.ServiceProposal {
	return dto_discovery.ServiceProposal{
		Id:          1,
		Format:      "service-proposal/v1",
		ServiceType: "openvpn",
		ServiceDefinition: dto.ServiceDefinition{
			Location:          nodeLocation,
			LocationOriginate: nodeLocation,
			SessionBandwidth:  dto.Bandwidth(10 * datasize.MB),
		},
		PaymentMethodType: dto.PAYMENT_METHOD_PER_TIME,
		PaymentMethod: dto.PaymentMethodPerTime{
			// 15 MYST/month = 0,5 MYST/day = 0,125 MYST/hour
			Price:    dto_discovery.Price{0.125, "MYST"},
			Duration: 1 * time.Hour,
		},
		ProviderId: nodeKey,
		ProviderContacts: []dto_discovery.Contact{
			{
				Type:       dto_discovery.CONTACT_NATS_V1,
				Definition: dto_discovery.ContactNATSV1{nodeKey},
			},
		},
	}
}
