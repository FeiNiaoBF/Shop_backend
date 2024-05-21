package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RotationReq struct {
	g.Meta `path:"/rotation/add" tags:"Rotation" method:"post" summary:"You first rotation api"`
	PicUrl string `json:"pic_url" v:"required#图片链接不能为空" dc:"图片链接"`
	Link   string `json:"link"    v:"required#跳转链接不能为空" dc:"跳转链接"`
	Sort   int    `json:"sort"    dc:"排序"`
}

type RotationRes struct {
	RotationId int `json:"rotation_id"`
}

type RotationDeleteReq struct {
	g.Meta `path:"/rotation/delete" method:"delete" tags:"轮播图" summary:"删除轮播图接口"`
	Id     uint `v:"min:1#请选择需要删除的轮播图" dc:"轮播图id"`
}

type RotationDeleteRes struct {
}
