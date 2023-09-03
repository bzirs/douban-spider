package main

import "github.com/bzirs/douban-spider/douban/login"

func main() {
	//
	//engine.Run(engine.Request{
	//	Url:        "https://movie.douban.com/chart",
	//	HandleFunc: movie.FetchMovieChart,
	//})

	login.Login()

}
