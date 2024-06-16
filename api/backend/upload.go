package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadToCloudReq struct {
	g.Meta `path:"/upload/cloud" method:"post" tag:"工具" mime:"multipart/form-data" dc:"上传文件到图床" summary:"上传文件到图床接口"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
}

type UploadToCloudRes struct {
	Url string `json:"url" dc:"图片地址"`
}
