package movie

import (
	"github.com/bzirs/douban-spider/douban/detail"
	"github.com/bzirs/douban-spider/douban/engine"
	"github.com/bzirs/douban-spider/utils"
	"regexp"
)

const RegMovie = `<a href="(https://movie.douban.com/subject/[\d]+/)"[^>]+>([^<]+)<span style="font-size:13px;">[^<]+</span>`

func FetchMovieChart(body []byte) engine.Result {

	reg := regexp.MustCompile(RegMovie)
	matches := reg.FindAllSubmatch(body, -1)

	result := engine.Result{}
	for _, match := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(match[1]),
			HandleFunc: detail.FetchDetail,
		})

		result.Contents = append(result.Contents, utils.FilterHan(string(match[2])))
	}
	return result

}
