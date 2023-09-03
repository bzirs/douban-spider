package login

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	Url       = "https://accounts.douban.com/j/mobile/login/basic"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
	Cookie    = `ll="118222"; bid=XugiDBoTSvg; _pk_id.100001.8cb4=6633ed01aca4f8fa.1684584399.; douban-fav-remind=1; push_noty_num=0; push_doumail_num=0; __utmv=30149280.23242; ct=y; ap_v=0,6.0; __utmc=30149280; ck=-2Ko; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1693731006%2C%22https%3A%2F%2Faccounts.douban.com%2F%22%5D; _pk_ses.100001.8cb4=1; __utma=30149280.702172822.1684584401.1693728828.1693731007.27; __utmz=30149280.1693731007.27.4.utmcsr=accounts.douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; __utmt=1; __utmb=30149280.30.10.1693731007`
)

func Login() {

	// payload
	data := url.Values{}

	data.Add("remember", "true")
	data.Add("name", "")
	data.Add("password", "")

	payload := strings.NewReader(data.Encode())

	request, err := http.NewRequest("POST", Url, payload)
	if err != nil {
		panic(err)
		return
	}

	request.Header.Set("Cookie", Cookie)
	request.Header.Set("User-Agent", UserAgent)
	request.Header.Set("Accept", "application/json, text/plain, */*")
	//request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "https://accounts.douban.com/passport/login_popup?login_source=anony")
	request.Header.Set("Origin", "https://accounts.douban.com")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
		return
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)

		return
	}

	fmt.Println(string(bytes))

}
