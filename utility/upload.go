package utility

import (
	"context"
	"log"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/storage"
)

func UploadToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	// 定义临时目录
	dirPath := "tmp/"
	filename, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}
	localPath := dirPath + filename
	// 取云端API
	bucket := getCfg(ctx, "qiniu.bucket")
	accessKey := getCfg(ctx, "qiniu.accessKey")
	secretKey := getCfg(ctx, "qiniu.secretKey")
	policy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := policy.UploadToken(mac)
	// 七牛云的配置
	qcfg := storage.Config{}
	qcfg.Zone = &storage.ZoneHuadongZheJiang2
	qcfg.UseHTTPS = true
	qcfg.UseCdnDomains = false
	//构建表单上传对象
	formUploader := storage.NewFormUploader(&qcfg)
	//上传结果的结构体
	ret := storage.PutRet{}

	//可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	//通过七牛云表单上传
	key := filename
	err = formUploader.PutFile(ctx, &ret, upToken, key, localPath, &putExtra)
	g.Dump(err)
	if err != nil {
		return "", err
	}
	// log
	log.Println(ret.Key, ret.Hash, ret.PersistentID)
	//删除本地临时文件
	err = os.RemoveAll(localPath)
	if err != nil {
		return "", err
	}
	//返回数据
	url = getCfg(ctx, "qiniu.url") + ret.Key
	return
}

func getCfg(ctx context.Context, path string) (url string) {
	return g.Cfg().MustGet(ctx, path).String()
}
