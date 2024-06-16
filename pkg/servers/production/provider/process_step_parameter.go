package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
)

type ProcessStepParameterProvider struct {
	proto.UnimplementedProcessStepParameterServer
}

func (u *ProcessStepParameterProvider) Query(ctx context.Context, in *proto.QueryProcessStepParameterRequest) (*proto.QueryProcessStepParameterResponse, error) {
	resp := &proto.QueryProcessStepParameterResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProcessStepParameter(in, resp, false)
	return resp, nil
}
