package biz

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"ratings/biz/model/db"
	"ratings/biz/model/model"
	"ratings/biz/pack"
	"ratings/kitex_gen/rating"
)

// RatingServiceImpl implements the last service interface defined in the IDL.
type RatingServiceImpl struct{}

// GetRating implements the RatingServiceImpl interface.
func (s *RatingServiceImpl) GetRating(ctx context.Context, req *rating.QueryRatingReq) (resp *rating.QueryRatingResp, err error) {
	// TODO: Your code here...
	resp = new(rating.QueryRatingResp)
	if req.PageSize == 0 || req.Page == 0 {
		resp.Code = rating.Code(codes.InvalidArgument)
		resp.Msg = "参数错误"
		return resp, nil
	}
	if req.GoodsID == 0 && req.ReviewID == 0 {
		resp.Code = rating.Code(codes.InvalidArgument)
		resp.Msg = "参数错误"
		return resp, nil
	}
	Query := db.Rating.WithContext(ctx)
	if req.GoodsID != 0 {
		Query = Query.Where(db.Rating.GoodsID.Eq(req.GoodsID))
	}
	if req.ReviewID != 0 {
		Query = Query.Where(db.Rating.ReviewID.Eq(req.ReviewID))
	}
	ratingModel := Query

	var total int64
	total, err = ratingModel.Count()
	if err != nil {
		resp.Code = rating.Code(codes.Internal)
		resp.Msg = "查询评分失败"
		return resp, nil
	}
	var ratings []*model.Rating
	if total > 0 {
		ratings, err = ratingModel.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find()
		if err != nil {
			resp.Code = rating.Code(codes.Internal)
			resp.Msg = "查询评分失败"
			return resp, nil
		}
	}
	resp.Code = rating.Code(codes.OK)
	resp.Msg = "查询评分成功"
	resp.Data = new(rating.Data)
	resp.Data.Total = total
	resp.Data.Rating = pack.Rating(ratings)

	return
}

// CreateRating implements the RatingServiceImpl interface.
func (s *RatingServiceImpl) CreateRating(ctx context.Context, req *rating.CreateRatingReq) (resp *rating.CreateRatingResp, err error) {
	// TODO: Your code here...
	resp = new(rating.CreateRatingResp)
	if req.GoodsID == 0 || req.Star == 0 || req.ReviewID == 0 {
		resp.Code = rating.Code(codes.InvalidArgument)
		resp.Msg = "参数错误"
		return resp, nil
	}
	err = db.Rating.WithContext(ctx).Create(
		&model.Rating{
			GoodsID:  req.GoodsID,
			Star:     req.Star,
			ReviewID: req.ReviewID,
		})
	if err != nil {
		resp.Code = rating.Code(codes.Internal)
		resp.Msg = "创建评分失败"
		return resp, nil
	}
	resp.Code = rating.Code(codes.OK)
	resp.Msg = "创建评分成功"
	return
}

// UpdateRating implements the RatingServiceImpl interface.
func (s *RatingServiceImpl) UpdateRating(ctx context.Context, req *rating.UpdateRatingReq) (resp *rating.UpdateRatingResp, err error) {
	// TODO: Your code here...
	resp = new(rating.UpdateRatingResp)
	if req.Id == 0 || req.Star == 0 || req.ReviewID == 0 || req.GoodsID == 0 {
		resp.Code = rating.Code(codes.InvalidArgument)
		resp.Msg = "参数错误"
		return resp, nil
	}

	_, err = db.Rating.WithContext(ctx).Where(db.Rating.ID.Eq(req.Id)).Updates(&model.Rating{
		ID:       req.Id,
		Star:     req.Star,
		ReviewID: req.ReviewID,
		GoodsID:  req.GoodsID,
	})
	if err != nil {
		resp.Code = rating.Code(codes.Internal)
		resp.Msg = "更新评分失败"
		return nil, err
	}

	resp.Code = rating.Code(codes.OK)
	resp.Msg = "更新评分成功"
	return
}

// DeleteRating implements the RatingServiceImpl interface.
func (s *RatingServiceImpl) DeleteRating(ctx context.Context, req *rating.DeleteRatingReq) (resp *rating.DeleteRatingResp, err error) {
	// TODO: Your code here...
	resp = new(rating.DeleteRatingResp)
	if req.Id == 0 {
		resp.Code = rating.Code(codes.InvalidArgument)
		resp.Msg = "参数错误"
		return resp, nil
	}
	_, err = db.Rating.WithContext(ctx).Where(db.Rating.ID.Eq(req.Id)).Delete()
	if err != nil {
		resp.Code = rating.Code(codes.Internal)
		resp.Msg = "删除评分失败"
		return nil, err
	}
	resp.Code = rating.Code(codes.OK)
	resp.Msg = "删除评分成功"
	return
}
