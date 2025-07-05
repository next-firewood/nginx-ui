package user

import (
	"context"
	"github.com/pkg/errors"
	"rui/internal/api"
	"rui/internal/api/user"
	"rui/internal/svc"
)

type UserDetailLogic struct {
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func NewUserDetailLogic(context context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Ctx:    context,
		SvcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *api.UuidForm) (resp *user.UserDetailResp, err error) {
	if req.Uuid == "" {
		return resp, errors.New("没有传入参数")
	}

	return &user.UserDetailResp{Uuid: req.Uuid, Name: "test"}, err
}
