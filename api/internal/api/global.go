package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

type PageInfo struct {
	Offset int `path:"offset"`
	Length int `path:"length"`
}

func PageInfoGet(c *gin.Context) (pageInfo PageInfo, err error) {
	if offset := c.Param("offset"); offset != "" {
		if pageInfo.Offset, err = strconv.Atoi(offset); err != nil {
			return pageInfo, errors.WithStack(err)
		}
	}

	if length := c.Param("length"); length != "" {
		if pageInfo.Length, err = strconv.Atoi(length); err != nil {
			return pageInfo, errors.WithStack(err)
		}
	}

	return pageInfo, err
}

type UuidForm struct {
	Uuid string `form:"uuid"`
}
