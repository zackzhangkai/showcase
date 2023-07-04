package rating

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"reviews/cmd/api/rpc/rating/kitex_gen/rating"
	"reviews/cmd/api/rpc/rating/kitex_gen/rating/ratingservice"
	"reviews/config"
)

var c ratingservice.Client

//var svc = "rating"
var svc = "ratings"

func init() {
	var err error
	//c, err = ratingservice.NewClient(svc, client.WithHostPorts("0.0.0.0:8888"))
	c, err = ratingservice.NewClient(svc, client.WithHostPorts(config.GetConfig().URL.Ratings))
	if err != nil {
		panic(err)
	}
}

func CreateRating(ctx context.Context, req *rating.CreateRatingReq) (*rating.CreateRatingResp, error) {
	resp, err := c.CreateRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetRating(ctx context.Context, req *rating.QueryRatingReq) (*rating.QueryRatingResp, error) {
	resp, err := c.GetRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateRating(ctx context.Context, req *rating.UpdateRatingReq) (*rating.UpdateRatingResp, error) {
	resp, err := c.UpdateRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteRating(ctx context.Context, req *rating.DeleteRatingReq) (*rating.DeleteRatingResp, error) {
	resp, err := c.DeleteRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
