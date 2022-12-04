package logic

import (
	"context"
	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	v, ok := l.ctx.Value("exp").(int64)
	if ok {
		l.Infof("exp: %d", v)
	}

	l.Infof("iat: %s", l.ctx.Value("iat"))
	l.Infof("userId: %s", l.ctx.Value("userId"))

	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	l.Infof("user: %v", user)

	return &types.OrderReply{
		Id:   user.GetId(),
		Name: user.GetName(),
	}, nil

}
