package main

import (
	"log"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/gin-gonic/gin"
)

var (
	serverUrl  string = "https://www.ptt.cc"
	baseUrl    string = "https://www.ptt.cc/bbs"
	kanban     string = "/Baseball"
	location   string = "/index.html"
	target     string = "大谷翔平"
	targetDate string = time.Now().String()
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	r.GET("/", func(c *gin.Context) {
		data := callAPI(baseUrl, kanban, location)
		if len(data) == 0 {
			data = "今日無" + target + "的文!"
		}
		html, _ := template.ParseFiles("static/index.html")
		date := getToday(targetDate)
		html.Execute(c.Writer, gin.H{
			"date": date,
			"data": data,
		})
	})
	r.Run()
}

func callAPI(base, kanban, location string) (ret string) {
	req := base + kanban + location
	soup.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
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
			if flag2 && (getElementDate(r, i) != getToday(targetDate)) {
				flag = false
			}
			if strings.Contains(element, target) == false {
				continue
			}
			tmp := strings.Split(element, `href="`)
			element = tmp[0] + `target="_blank" ` + `href="` + serverUrl + tmp[1]
			ret += element
		}
		flag2 = true
		req = getPreUrl(doc)
	}
	return
}

func getToday(now string) string {
	tmp := strings.Split(now, ` `)
	ret := tmp[0]
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
	tmp := strings.Split(preUrl, `href="`)
	tmp2 := strings.Split(tmp[1], `"`)
	return serverUrl + tmp2[0]
}
