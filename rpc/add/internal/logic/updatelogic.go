package logic

import (
	"bookstore/rpc/model"
	"context"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *add.AddReq) (*add.AddResp, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.Model.Update(model.Book{
		Book:  in.Book,
		Price: in.Price,
	})

	if err != nil {
		return nil, err
	}

	return &add.AddResp{
		Ok: true,
	}, nil
}
