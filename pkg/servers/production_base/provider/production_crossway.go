package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
)

type ProductionCrosswayProvider struct {
	proto.UnimplementedProductionCrosswayServer
}

func (u *ProductionCrosswayProvider) Query(ctx context.Context, in *proto.QueryProductionCrosswayRequest) (*proto.QueryProductionCrosswayResponse, error) {
	resp := &proto.QueryProductionCrosswayResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProductionCrossway(in, resp, false)
	return resp, nil
}
