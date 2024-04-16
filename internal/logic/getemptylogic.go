package logic

import (
	"context"

	"demo2/internal/svc"
	"demo2/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmptyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmptyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmptyLogic {
	return &GetEmptyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmptyLogic) GetEmpty(in *proto.EmptyReq) (*proto.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &proto.EmptyResp{}, nil
}
