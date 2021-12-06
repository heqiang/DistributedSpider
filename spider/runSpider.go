package spider

import (
	"ces.mod/app"
	"fmt"
	"reflect"
)



func  Run()  {
	//url := "https://www.qiushibaike.com/text/"
	//var iteminfo []parsespider.ItemInfo
	//var qiushi  parsespider.QiuShi
	allSpider :=app.GetAllSpider()
	for _,value :=range allSpider{
		 for _,v :=range value{
			 spider := reflect.ValueOf(v)
			 fmt.Println(spider.NumMethod())
			 //infoFunc :=spider.MethodByName("GetList")
			 //infoFunc.Call( []reflect.Value{})
		 }

	}
	//byteV,err := qiushi.GetResponse(url)
	//UrlList, err := qiushi.GetList(byteV)
	//utils.HandError(err)
	//for _,url := range UrlList{
	//	byteValue,err := qiushi.GetResponse(url)
	//	utils.HandError(err)
	//	tile, err := qiushi.GetTile(byteValue)
	//	utils.HandError(err)
	//	author, err := qiushi.GetAuthor(byteValue)
	//	utils.HandError(err)
	//	putime, err := qiushi.GetPubTime(byteValue)
	//	utils.HandError(err)
	//	content, err := qiushi.GetContent(byteValue)
	//	utils.HandError(err)
	//	item:=parsespider.ItemInfo{
	//		Url: url,
	//		Title:tile,
	//		Author: author,
	//		PubTime: putime,
	//		Content: content,
	//	}
	//	iteminfo = append(iteminfo,item)
	//
	//}
	//for _,key :=range iteminfo{
	//	fmt.Printf("url:%s,时间:%s\n",key.Url,key.PubTime)
	//}
}