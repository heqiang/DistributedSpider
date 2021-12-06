package parsespider

import (
	"ces.mod/utils"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type  QiuShi struct {
	WebName string `json:"webname"`
	SpiderName string `json:"spidername"`
	StartUrl   []string 	`json:"starturl"`
}


func (q *QiuShi)GetList(resbyter []byte) ([]string, error)  {
	var urlList []string
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(resbyter)))
	utils.HandError(err)
	dom.Find(`div.col1>div>a`).Each(func(i int, selection *goquery.Selection) {
			url,ok := selection.Attr("href")
			if ok{
				realUrl := "https://www.qiushibaike.com"+url
				urlList = append(urlList,realUrl)
			}
	})
	return urlList,nil
}
func (q *QiuShi)GetTile(resByte []byte) (string, error)  {
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(resByte)))
	utils.HandError(err)
	title :=dom.Find(`h1.article-title`).Text()
	return strings.TrimSpace(title),nil
}
func (q *QiuShi)GetPubTime(resByte []byte) (time.Time, error)  {
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(resByte)))
	utils.HandError(err)
	pub :=strings.TrimSpace(dom.Find(`span.stats-time`).Text())
	time_,err:= time.Parse("2006-01-02 15:04",pub)
	utils.HandError(err)
	return time_,nil
}
func (q *QiuShi)GetAuthor(resByte []byte) (string,error)  {
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(resByte)))
	utils.HandError(err)
	author :=dom.Find(`span.side-user-name`).Text()
	return strings.TrimSpace(author),nil
}
func (q *QiuShi)GetContent(resByte []byte) (string,error)  {
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(string(resByte)))
	utils.HandError(err)
	content :=dom.Find(`div.content`).Text()
	return content,nil
}
func (q *QiuShi)GetResponse() ([]byte,error)  {
	for _,url :=range  q.StartUrl{
		client := &http.Client{}
		req, err := http.NewRequest("GET",url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		return  bodyText,nil
	}
	return nil,nil
}