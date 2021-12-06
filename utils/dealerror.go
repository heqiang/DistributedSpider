package utils

import "fmt"

func HandError(err error)  {
	if err!=nil{
		fmt.Println(err)
		return
	}
}
