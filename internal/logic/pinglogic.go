package logic

import (
	"context"

	"demo2/internal/svc"
	"demo2/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *proto.Request) (*proto.Response, error) {
	// todo: add your logic here and delete this line

	return &proto.Response{}, nil
}
