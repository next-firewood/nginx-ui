package repo

import (
	"context"
	"github.com/pkg/errors"
	"rui/common/database"
	"rui/common/errorx"
)

type User struct {
	database.GormModel
	Username string `gorm:"uniqueIndex;size:50;comment:用户名"`
	Password string `gorm:"size:50;comment:密码"`
	Level    int32  `gorm:"size:1;comment:等级 1-最高级"`
}

func (User) TableName() string {
	return "user"
}

type UserDber interface {
	UserDetail(ctx context.Context, req User) (res User, err error)
	UserNew(ctx context.Context, req User) (res User, err error)
}

func (r *globalDb) UserDetail(ctx context.Context, req User) (res User, err error) {
	cd := r.DB.WithContext(ctx)

	if req.Id > 0 {
		cd = cd.Where("id = ?", req.Id)
	}

	if req.Uuid != "" {
		cd = cd.Where("uuid = ?", req.Uuid)
	}

	if req.Username != "" {
		cd = cd.Where("username = ?", req.Username)
	}

	if req.Password != "" {
		cd = cd.Where("password = ?", req.Password)
	}

	if req.Level > 0 {
		cd = cd.Where("level = ?", req.Level)
	}

	if err = cd.Limit(1).Find(&res).Error; err != nil {
		return res, errors.WithStack(err)
	}

	return res, err
}

func (r *globalDb) UserNew(ctx context.Context, req User) (res User, err error) {
	isDup, err := errorx.SqlDupErr(r.GetDB(ctx).Create(&req).Error)
	if isDup {
		return res, err
	}

	return req, errors.WithStack(err)
}
