package main

import "fmt"

type Metric struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}


// demo_select6.go
func main() {
	//github修改第一次 ，下次本地修改
	//本地修改
	//github修改第二次 
	ch := make(chan int, 1)

	// 发送1个数据关闭channel
	ch <- 1
	//close(ch)
	// 不停读数据直到channel没有有效数据
	for {
		select {
		case v, ok := <-ch:
			fmt.Println("v:", v, ", ok:", ok )
			if !ok {
				fmt.Println(v,ok)
			}
			if v==2 {
	     		break
			}
		default:
			ch <-2
		}


	}
}

