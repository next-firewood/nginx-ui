package repo

import (
	"context"
	"github.com/pkg/errors"
	"rui/common/errorx"
)

type SysConf struct {
	Id      uint64 `gorm:"primarykey" json:"id"`
	Code    string `gorm:"size:50;uniqueIndex:sys_conf;comment:编码"`
	Node    string `gorm:"size:50;uniqueIndex:sys_conf;comment:节点"`
	Value   string `gorm:"comment:值"`
	IsLocal int32  `gorm:"size:1;default:2;comment:是否是本地"`
}

func (SysConf) TableName() string {
	return "system_conf"
}

type SysConfDber interface {
	SysConfDetailByNode(ctx context.Context, req SysConf) (res SysConf, err error)
	SysConfDetailNew(ctx context.Context, req SysConf) (err error)
}

func (r *globalDb) SysConfDetailByNode(ctx context.Context, req SysConf) (res SysConf, err error) {
	cd := r.DB.WithContext(ctx)

	if req.Node != "" {
		cd = cd.Where("node = ?", req.Node)
	}

	if req.Code != "" {
		cd = cd.Where("code = ?", req.Code)
	}

	if req.IsLocal > 0 {
		cd = cd.Where("is_local = ?", req.IsLocal)
	}

	if req.Id > 0 {
		cd = cd.Where("id = ?", req.Id)
	}

	if err = cd.Limit(1).Find(&res).Error; err != nil {
		return res, errors.WithStack(err)
	}

	return res, err
}

func (r *globalDb) SysConfDetailNew(ctx context.Context, req SysConf) (err error) {
	isDup, err := errorx.SqlDupErr(r.GetDB(ctx).Create(&req).Error)
	if isDup {
		return err
	}

	return errors.WithStack(err)
}
