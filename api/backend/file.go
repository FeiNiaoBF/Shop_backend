package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mime:"multipart/form-data" tags:"文件" dc:"上传文件" summary:"上传本地文件接口"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}
type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"图片地址"`
}
