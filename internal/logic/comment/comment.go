package comment

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"goBack/internal/consts"
	"goBack/internal/dao"
	"goBack/internal/model"
	"goBack/internal/service"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

func (s *sComment) AddComment(ctx context.Context, in model.CommentAddInput) (res *model.CommentAddOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.CommentAddOutput{}, err
	}
	return &model.CommentAddOutput{Id: gconv.Uint(id)}, nil
}

// 兼容处理：优先根据收藏id删除，收藏id为0；再根据对象id和type删除
func (s *sComment) DeleteComment(ctx context.Context, in model.CommentDeleteInput) (res *model.CommentDeleteOutput, err error) {
	//优先根据收藏id删除
	if in.Id != 0 {
		_, err = dao.CommentInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.CommentDeleteOutput{Id: gconv.Uint(in.Id)}, nil
	} else {
		//	收藏id为0；再根据对象id和type删除
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CommentInfo.Ctx(ctx).OmitEmpty(). //注意：需要过滤空值
									Where(in).Delete()
		if err != nil {
			return &model.CommentDeleteOutput{}, err
		}
		return &model.CommentDeleteOutput{Id: gconv.Uint(id)}, nil
	}
}

// GetList 查询内容列表
func (s *sComment) GetList(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error) {
	m := dao.CommentInfo.Ctx(ctx)
	out = &model.CommentGetListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CommentListItem{}, //数据为空时返回空数组 而不是null
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)

	// 条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.CommentInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count 而不是像之前一样查sql结果赋值到结构体中
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, err
	}
	//进一步优化：只查询相关的模型关联
	if in.Type == consts.CommentTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else if in.Type == consts.CommentTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, err
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}

func CommentCount(ctx context.Context, objectId uint, collectionType uint8) (count int, err error) {
	condition := g.Map{
		dao.CommentInfo.Columns().ObjectId: objectId,
		dao.CommentInfo.Columns().Type:     collectionType,
	}
	count, err = dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return 0, err
	}
	return
}

func CheckIsComment(ctx context.Context, in model.CheckIsCollectInput) (bool, error) {
	condition := g.Map{
		dao.CommentInfo.Columns().UserId:   ctx.Value(consts.CtxUserId),
		dao.CommentInfo.Columns().ObjectId: in.ObjectId,
		dao.CommentInfo.Columns().Type:     in.Type,
	}
	count, err := dao.CommentInfo.Ctx(ctx).Where(condition).Count()
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
