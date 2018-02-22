package main
/*
    功能：消耗CPU，用于压力测试
	(TODO：通过http开启或暂停测试)
*/

import (
	"time"
	"math"
	"fmt"
)

func main(){
	fmt.Println("****** Heat Cpu v0.1******")
	fmt.Println("Ready...")
	time.Sleep(time.Second * 5)
	fmt.Println("Go!!!")

	for{
		go func()(){
			_ = math.Sqrt(123456789.123456789)
		}()
	}

	fmt.Println("Done")
}
