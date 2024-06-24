package collection

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/internal/consts"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/model/entity"
	"goBack/internal/service"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (s *sCollection) AddCollection(ctx context.Context, in model.CollectionAddInput) (res *model.CollectionAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.CollectionAddOutput{}, err
	}
	return &model.CollectionAddOutput{Id: gconv.Uint(id)}, nil
}

// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
func (s *sCollection) DeleteCollection(ctx context.Context, in model.CollectionDeleteInput) (res *model.CollectionDeleteOutput, err error) {
	//优先根据收藏id删除
	if in.Id != 0 {
		_, err = dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.CollectionDeleteOutput{Id: gconv.Uint(in.Id)}, nil
	} else {
		//	收藏id为0；再根据对象id和type删除
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty(). //注意：需要过滤空值
									Where(in).Delete()
		if err != nil {
			return &model.CollectionDeleteOutput{}, err
		}
		return &model.CollectionDeleteOutput{Id: gconv.Uint(id)}, nil
	}
}

// GetList 查询内容列表
func (s *sCollection) GetList(ctx context.Context, in model.CollectionGetListInput) (out *model.CollectionGetListOutput, err error) {
	m := dao.CollectionInfo.Ctx(ctx)
	out = &model.CollectionGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 执行查询
	var list []*entity.CollectionInfo
	if err = listModel.WithAll().Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	// 计数
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	//不指定item的键名用：Scan
	if err := listModel.WithAll().Scan(&out.List); err != nil {
		return out, err
	}
	return
}
