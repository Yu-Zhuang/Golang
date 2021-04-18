package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
)

var (
	serverUrl string = "https://www.ptt.cc"
	baseUrl   string = "https://www.ptt.cc/bbs"
	//kanban    string = "/Baseball"
	location string = "/index.html"
	//target    string = "大谷翔平"
	//date string = time.Now().String()
)

type requestBody struct {
	Kanban string `json:"kanban" form:"kanban"`
	Date   string `json:"date" form:"date"`
	Target string `json:"target" form:"target"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/api/getResult", func(c *gin.Context) {
		var reqBody requestBody
		c.Bind(&reqBody)
		if msg, ok := isInputOK(reqBody); ok == false {
			c.JSON(http.StatusBadRequest, gin.H{
				"data": msg,
			})
		}
		reqBody.Kanban = "/" + reqBody.Kanban
		re, _ := regexp.Compile(`^(.+)/(.+)/(.+)$`)
		reqBody.Date = re.ReplaceAllString(reqBody.Date, `$1-$2-$3`)
		fmt.Println(reqBody.Kanban, reqBody.Date, reqBody.Target)
		data := callAPI(baseUrl, reqBody.Kanban, reqBody.Date, reqBody.Target)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	r.Run()
}

func isInputOK(input requestBody) (msg string, ok bool) {
	if len(input.Date) == 0 || len(input.Kanban) == 0 || len(input.Target) == 0 {
		return "不能有空值", false
	}
	return "ok", true
}

func callAPI(base, kanban, date, target string) (ret string) {
	req := base + kanban + location
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		"cookie":     "over18=1",
	}
	flag := true
	flag2 := false
	for flag {
		source, err := soup.Get(req)
		if err != nil {
			log.Fatal(err)
		}
		doc := soup.HTMLParse(source)
		r := doc.FindAll("div", "class", "r-ent")
		for i := range r {
			element := r[i].Find("div", "class", "title").HTML()
			if flag2 && (getElementDate(r, i) != getToday(date)) {
				flag = false
			}
			if strings.Contains(element, target) == false {
				continue
			}
			if strings.Contains(element, `href="`) == true {
				tmp := strings.Split(element, `href="`)
				element = tmp[0] + `target="_blank" ` + `href="` + serverUrl + tmp[1]
			}
			ret += element
		}
		flag2 = true
		req = getPreUrl(doc)
	}
	return
}

func getToday(now string) string {
	ret := now
	if strings.Contains(now, ` `) == true {
		tmp := strings.Split(now, ` `)
		ret = tmp[0]
	}
	re, _ := regexp.Compile(`^(.+)-(.+)-(.+)$`)
	ret = re.ReplaceAllString(ret, `$2/$3`)
	if string(ret[0]) == `0` {
		ret = ret[1:]
	}
	return ret
}

func getElementDate(r []soup.Root, i int) string {
	tmp := r[i].Find("div", "class", "meta").Find("div", "class", "date").HTML()
	tmp = strings.ReplaceAll(tmp, `<div class="date">`, "")
	tmp = strings.ReplaceAll(tmp, `</div>`, "")
	ret := strings.ReplaceAll(tmp, ` `, ``)
	return ret
}

func getPreUrl(doc soup.Root) string {
	preUrlList := doc.FindAll("a")
	var preUrl string
	for i := range preUrlList {
		if strings.Contains(preUrlList[i].HTML(), "上頁") {
			preUrl = preUrlList[i].HTML()
			break
		}
	}
	var tmp2 []string
	if strings.Contains(preUrl, `href="`) == true {
		tmp := strings.Split(preUrl, `href="`)
		tmp2 = strings.Split(tmp[1], `"`)
	}
	if len(tmp2) == 0 {
		return serverUrl
	}
	return serverUrl + tmp2[0]
}
