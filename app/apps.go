package app

import (
	"ces.mod/spider/parsespider"
)

var AllSpider = []map[int]interface{}{
	{
		1:&parsespider.QiuShi{
			WebName: "aa",
			SpiderName: "ff",
			StartUrl: "https://www.qiushibaike.com/text",
		},
	},
}

func GetAllSpider() []map[int]interface{}  {
	return AllSpider
}