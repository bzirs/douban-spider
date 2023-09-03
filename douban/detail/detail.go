package detail

import (
	"github.com/bzirs/douban-spider/douban/engine"
	"regexp"
)

const RegName = `<span property="v:itemreviewed">([^>]+)</span>`
const RegDate = `<span property="v:initialReleaseDate"[^>]+>([^>]+)</span>`
const RegIntroduce = `<span property="v:summary".*>([^<]+)</span>`

func FetchDetail(body []byte) engine.Result {
	detail := Detail{}

	detail.Name = extractString(body, RegName)
	detail.Date = extractString(body, RegDate)
	detail.Introduce = extractString(body, RegIntroduce)

	return engine.Result{
		Contents: []interface{}{detail},
	}

}

func extractString(body []byte, regRule string) string {
	reg := regexp.MustCompile(regRule)
	match := reg.FindSubmatch(body)
	if match != nil {
		return string(match[1])

	} else {
		return ""
	}

}
