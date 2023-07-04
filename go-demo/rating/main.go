package main

import (
	"log"
	"ratings/biz"
	_ "ratings/biz/dal"
	rating "ratings/kitex_gen/rating/ratingservice"
)

func main() {
	svr := rating.NewServer(new(biz.RatingServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
