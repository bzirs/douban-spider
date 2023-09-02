package engine

import (
	"github.com/bzirs/douban-spider/utils"
	"log"
)

func Run(seed ...Request) {
	requests := append([]Request{}, seed...)

	for len(requests) > 0 {

		req := requests[0]

		requests = requests[1:]

		body, err := utils.Fetch(req.Url)
		if err != nil {
			log.Printf("Fetch Error: %s", err)
			continue

		}

		result := req.HandleFunc(body)

		requests = append(requests, result.Requests...)

		for _, content := range result.Contents {
			log.Printf("Current Content: %s", content)
		}

	}
}
