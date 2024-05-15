package provider

import (
	"context"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProcessStepMatchRuleProvider struct {
	proto.ProcessStepMatchRuleServer
}

func (u *ProcessStepMatchRuleProvider) Query(ctx context.Context, in *proto.QueryProcessStepMatchRuleRequest) (*proto.QueryProcessStepMatchRuleResponse, error) {
	resp := &proto.QueryProcessStepMatchRuleResponse{
		Code: proto.Code_Success,
	}
	logic.QueryProcessStepMatchRule(in, resp, false)
	return resp, nil
}
