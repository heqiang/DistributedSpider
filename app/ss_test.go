package app

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_dd(t *testing.T)  {
	GetAllSpider()
	for _,value :=range AllSpider{
		for _,v :=range value{
			 info := reflect.ValueOf(v)
			 //funcs :=
			 fmt.Println(info)

		}
	}
}
