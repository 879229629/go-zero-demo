package logic

import (
	"context"
	"fmt"

	"go-zero-demo/mall/user/rpc/internal/svc"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	l.Infof("rpc req: %v", in)
	return &user.UserResponse{Id: in.GetId(), Name: fmt.Sprintf("rpc name: %s", in.GetId())}, nil
}
