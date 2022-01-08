// ## 请在下方进行输入 ( 支持Markdown、插入公式及 LaTex 数学公式，点击上方按钮“环境说明”查看详情 )

//两个协程交替打印1-100的奇数偶数
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch1, ch2 := make(chan struct{}, 1), make(chan struct{})

	go func() {
		defer wg.Done()
		for i := 1; i < 100; i += 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- struct{}{}
		}
	}()

	ch1 <- struct{}{}

	wg.Wait()
}
