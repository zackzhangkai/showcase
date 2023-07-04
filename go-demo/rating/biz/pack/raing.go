package pack

import (
	"ratings/biz/model/model"
	"ratings/kitex_gen/rating"
)

// Users Convert model.Review list to api.Review list
func Rating(models []*model.Rating) []*rating.Rating {
	reviews := make([]*rating.Rating, 0, len(models))
	for _, m := range models {
		if u := Review(m); u != nil {
			reviews = append(reviews, u)
		}
	}
	return reviews
}

// Review Convert model.Review to api.Review
func Review(model *model.Rating) *rating.Rating {
	if model == nil {
		return nil
	}
	return &rating.Rating{
		ReviewID: model.ReviewID,
		GoodsID:  model.GoodsID,
		Star:     model.Star,
		Average:  model.Average,
	}
}
