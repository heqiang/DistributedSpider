package parsespider

import "time"

type BaseInterface interface {
	GetList([]byte) ([]string, error)
	GetTile([]byte) (string, error)
	GetPubTime([]byte) (time.Time, error)
	GetAuthor([]byte) (string,error)
	GetContent([]byte) (string,error)
	GetResponse(url string) ([]byte,error)
}


type  ItemInfo struct {
	Url  string `json:"url"`
	Title string`json:"title"`
	PubTime time.Time `json:"pubtime"`
	Author string `json:"author"`
	Content string `json:"content"`
}