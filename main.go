package main

import (
	"github.com/bzirs/douban-spider/douban/engine"
	"github.com/bzirs/douban-spider/douban/movie"
)

func main() {

	engine.Run(engine.Request{
		Url:        "https://movie.douban.com/chart",
		HandleFunc: movie.FetchMovieChart,
	})

}
