package controller

import (
	"context"
	"goBack/api/backend"
	"goBack/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var Upload cUpload

type cUpload struct {
}

func (c *cUpload) UploadToCloud(ctx context.Context, req *backend.UploadToCloudReq) (res *backend.UploadToCloudRes, err error) {
	if req.File == nil {
		err = gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
		return
	}
	// log.Println(req.File)
	res = new(backend.UploadToCloudRes)
	url, err := utility.UploadToCloud(ctx, req.File)
	if err != nil {
		err = gerror.NewCode(gcode.CodeInternalError, err.Error())
		return
	}
	res.Url = url
	return
}
