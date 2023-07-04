package pack

import (
	"reviews/biz/model/model"
	"reviews/biz/model/review"
)

// Users Convert model.Review list to api.Review list
func Reviews(models []*model.Review) []*review.Review {
	reviews := make([]*review.Review, 0, len(models))
	for _, m := range models {
		if u := Review(m); u != nil {
			reviews = append(reviews, u)
		}
	}
	return reviews
}

// Review Convert model.Review to api.Review
func Review(model *model.Review) *review.Review {
	if model == nil {
		return nil
	}
	return &review.Review{
		ID:      model.ID,
		Name:    model.Name,
		GoodsID: model.GoodsID,
		Content: model.Content,
		Author:  model.Author,
	}
}
